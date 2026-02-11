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
				um.user_id = ? AND pr.rule = 'OWN'
		);`
	return query
}

func UserInMarket() string {
	query :=
		`SELECT EXISTS (
			SELECT 1
			FROM 
				users_markets um
			WHERE
				um.user_id = ? AND um.market_id = ?
		);`
	return query
}
