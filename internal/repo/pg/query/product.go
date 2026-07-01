package query

func GetProduct() string {
	query :=
		`SELECT
			p.product_id,
			p.name,
			p.description,
			p.category,
			p.price,
			p.available_stock,
			p.reserved_stock,
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
			p.available_stock,
			p.reserved_stock,
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

func GetProductsByIds() string {
	query :=
		`SELECT
			p.product_id,
			p.name,
			p.description,
			p.category,
			p.price,
			p.available_stock,
			p.reserved_stock,
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
			p.product_id IN (?);`
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
			p.available_stock,
			p.reserved_stock,
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
				available_stock,
				reserved_stock,
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
				0,
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
			p.available_stock = ?,
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
			price = ?
		WHERE
			product_id = ?;`
	return query
}

func DecreaseProductCountFromOrder() string {
	query :=
		`WITH decrease AS (
			VALUES ?
			)
		UPDATE products p
		SET
			count = p.count - d.column2
		FROM 
			decrease d(id, decrement)
		WHERE
			p.id = u.id;`
	return query
}

func IncreaseProductCountFromOrder() string {
	query :=
		`WITH increase AS (
			VALUES ?
			)
		UPDATE products p
		SET
			count = p.count + d.column2
		FROM 
			increase i(id, increment)
		WHERE
			p.id = u.id;`
	return query
}

func ReserveStock() string {
	return `UPDATE products
				SET available_stock = available_stock - ?,
					reserved_stock = reserved_stock + ?,
				WHERE product_id = ?
				AND available_stock >= ?;`
}

func ConfirmReserve() string {
	return `UPDATE products
				SET reserved_stock = reserved_stock - ?,
				WHERE product_id = ?;`
}

func DeclineReserve() string {
	return `UPDATE products
				SET available_stock = available_stock + ?,
					reserved_stock = reserved_stock - ?,
				WHERE product_id = ?;`
}
