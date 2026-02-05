package query

func GetProduct() string {
	query :=
		`SELECT
			p.product_id,
			p.name,
			p.description,
			p.category,
			p.price,
			p.count,
			p.active,
			p.options,
			p.article,
			p.inserted_by,
			p.inserted,
			p.updated_by,
			p.updated
		FROM 
			products p`
	return query
}

func GetProductById() string {
	query :=
		`SELECT
			p.product_id,
			p.name,
			p.description,
			p.category,
			p.price,
			p.count,
			p.active,
			p.options,
			p.article,
			p.inserted_by,
			p.inserted,
			p.updated_by,
			p.updated
		FROM 
			products p
		WHERE
			p.product_id = ?`
	return query
}

func GetProductByArticle() string {
	query :=
		`SELECT
			p.product_id,
			p.name,
			p.description,
			p.category,
			p.price,
			p.count,
			p.active,
			p.options,
			p.article,
			p.inserted_by,
			p.inserted,
			p.updated_by,
			p.updated
		FROM 
			products p
		WHERE
			p.article = ?`
	return query
}
