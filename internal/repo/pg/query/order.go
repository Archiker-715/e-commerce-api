package query

func CreateOrder() string {
	query :=
		`INSERT INTO orders(
			order_id,
			user_id,
			order_price,
			products,
			paid_expired,
			paid,
			inserted_by,
			inserted
			)
			SELECT 
				?,
				?,
				?,
				?
				?,
				?,
				?,
				?
				;`
	return query
}

func GetOrderById() string {
	query :=
		`SELECT
			o.order_id,
			o.user_id,
			o.order_price,
			o.products,
			o.paid_expired,
			o.paid,
			o.inserted_by,
			o.inserted,
			o.updated_by,
			o.updated
		FROM 
			orders o
		WHERE
			o.order_id = ?;`
	return query
}

func MarkExpired() string {
	query :=
		`UPDATE orders
		SET
			paid_expired = true
		WHERE
			order_id = ?;`
	return query
}

func MarkPaid() string {
	query :=
		`UPDATE orders
		SET
			paid = true
		WHERE
			order_id = ?;`
	return query
}
