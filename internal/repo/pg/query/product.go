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
			products p;`
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
			p.product_id = ?;`
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
			p.article = ?;`
	return query
}

func CreateProduct() string {
	query :=
		`WITH newProduct AS (
			INSERT INTO products(
				name,
				description,
				category,
				price,
				count,
				active,
				options,
				article,
				inserted_by,
				inserted
			)
			SELECT 
				?,
				?,
				?,
				?,
				?,
				?,
				?,
				?,
				?,
				?
			RETURNING 
				product_id, inserted_by, inserted
		)
		INSERT INTO product_rules (
			product_id,
			market_id,
			rule,
			inserted_by,
			inserted
		)
		SELECT
			product_id,
			?,
			'OWN',
			inserted_by,
			inserted
			;`
	return query
}

func UpdateProduct() string {
	query :=
		`UPDATE products
		SET
			p.name = ?,
			p.description = ?,
			p.category = ?,
			p.price = ?,
			p.count = ?,
			p.active = ?,
			p.options = ?,
			p.updated_by = ?,
			p.updated = ?
		WHERE
			product_id = ?;`
	return query
}

func DeleteProduct() string {
	query :=
		`DELETE FROM products
		WHERE
			product_id = ?;`
	return query
}

func UpdatePrice() string {
	query :=
		`UPDATE products
		SET
			p.price = ?
		WHERE
			product_id = ?;`
	return query
}
