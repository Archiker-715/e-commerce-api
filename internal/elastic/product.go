package elastic

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/Archiker-715/e-commerce-api/internal/entity"
)

func (e *elastic) SearchProduct(indexName, keyword string) error {
	var buf strings.Builder
	query := fmt.Sprintf(`{
        "query": {
            "match": {
                "content": "%s"
            }
        }
    }`, keyword)
	buf.WriteString(query)

	res, err := e.es.Search(
		e.es.Search.WithContext(context.Background()),
		e.es.Search.WithIndex(indexName),
		e.es.Search.WithBody(strings.NewReader(buf.String())),
		e.es.Search.WithTrackTotalHits(true),
	)
	if err != nil {
		return fmt.Errorf("err when search in es: %w", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		body, _ := io.ReadAll(res.Body)
		return fmt.Errorf("res from es err when create idx: %v", body)
	}

	var r map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		return fmt.Errorf("parsing es body err: %w", err)
	}

	hits := r["hits"].(map[string]interface{})["hits"].([]interface{})
	for _, hit := range hits {
		doc := hit.(map[string]interface{})
		source := doc["_source"]
		fmt.Printf("Найден документ: %+v\n", source)
	}
	return nil
}

func (e *elastic) BulkIndexProduct(indexName string, docs []entity.Product) error {

	var buf bytes.Buffer
	for _, doc := range docs {
		meta := fmt.Sprintf(`{ "index" : { "_index" : "%s" } }%s`, indexName, "\n")
		buf.WriteString(meta)

		data, err := json.Marshal(doc)
		if err != nil {
			return fmt.Errorf("marshal doc err: %w", err)
		}
		buf.Write(data)
		buf.WriteByte('\n')
	}

	res, err := e.es.Bulk(
		bytes.NewReader(buf.Bytes()),
		e.es.Bulk.WithContext(context.Background()),
	)
	if err != nil {
		return fmt.Errorf("bulk err: %w", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		body, _ := io.ReadAll(res.Body)
		return fmt.Errorf("res from es err when bulk docs: %v", body)
	}
	return nil
}
