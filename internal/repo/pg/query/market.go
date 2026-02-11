package query

func AddMarket() string {
	query :=
		`WITH newMarket AS (
			INSERT INTO markets(
				market_name,
				inserted_by,
				inserted
				)
			SELECT 
				?,
				?,
				?
			RETURNING market_id, inserted_by, inserted
			)
		INSERT INTO users_markets(
			user_id,
			market_id,
			market_owner_user_id,
			inserted_by,
			inserted
			)
			SELECT 
				?,
				market_id,
				inserted_by,
				inserted_by,
				inserted
				;`
	return query
}

func LinkUserMarket() string {
	query :=
		`INSERT INTO users_markets(
			user_id,
			market_id,
			inserted_by,
			inserted
			)
			SELECT 
				?,
				?,
				?,
				?
				;`
	return query
}

func CheckOwner() string {
	query :=
		`SELECT EXISTS (
			SELECT 1
			FROM 
				users_markets um
			WHERE
				um.market_id = ? AND market_owner_user_id = ?
		);`
	return query
}
