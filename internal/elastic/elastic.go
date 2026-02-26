package elastic

import (
	"io"
	"log"

	"github.com/elastic/go-elasticsearch/v8"
)

type elastic struct {
	es *elasticsearch.Client
}

func NewElastic() *elastic {
	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf("new es client err: %v\n", err)
	}

	res, err := es.Info()
	if err != nil {
		log.Fatalf("resp get err: %v\n", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		body, _ := io.ReadAll(res.Body)
		log.Fatalf("res from es err: %v\n", body)
	}

	return &elastic{
		es: es,
	}
}

func (e *elastic) newIndex(indexName string) {
	res, err := e.es.Indices.Create(indexName)
	if err != nil {
		log.Fatalf("cannot create index: %v\n", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		body, _ := io.ReadAll(res.Body)
		log.Fatalf("res from es err when create idx: %v\n", body)
	}
}
