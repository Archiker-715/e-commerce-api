package query

func AddMarket() string {
	query :=
		`WITH newMarket AS (
			INSERT INTO markets(
				market_id,
				market_name,
				inserted_by,
				inserted
				)
			SELECT 
				?,
				?,
				?,
				?
			RETURNING market_id, inserted_by, inserted
			)
		INSERT INTO users_markets(
			user_id,
			market_id,
			inserted_by,
			inserted
			)
			SELECT 
				?,
				market_id,
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
				inserted_by,
				inserted
				;`
	return query
}
