package query

func CheckPermission() string {
	query :=
		`SELECT EXISTS (
			SELECT 1
			FROM 
				product_rules pr
			JOIN 
				users_markets um ON pr.market_id = um.market_id
			WHERE
				um.user_id = ? AND pr.rule = ?
		);`
	return query
}

func AddPermission() string {
	query :=
		`INSERT INTO product_rules(
			product_id,
			market_id,
			rule,
			inserted,
			inserted_by
		)
		SELECT 
			?,
			?,
			?,
			?,
			?;`
	return query
}
