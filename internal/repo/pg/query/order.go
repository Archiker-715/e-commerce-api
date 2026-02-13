package query

func CreateOrder() string {
	query :=
		`INSERT INTO orders(
			order_id,
			user_id,
			order_price,
			products,
			temp,
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
				?
				;`
	return query
}
