package query

func GetUserCart() string {
	query :=
		`SELECT
			uc.name,
			uc.price,
			uc.count,
		FROM 
			users_cart uc
		WHERE
			uc.user_id = ?;`
	return query
}
