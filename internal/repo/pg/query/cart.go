package query

func GetUserCart() string {
	query :=
		`SELECT
			uc.name,
			uc.total_price,
			uc.count,
		FROM 
			users_cart uc
		WHERE
			uc.user_id = ?;`
	return query
}

func AddProductToCart() string {
	query :=
		`INSERT INTO users_cart(
			user_id,
			product_id,
			name,
			count,
			unit_price,
			total_price
		)
		SELECT 
			?,
			product_id,
			name,
			1,
			unit_price,
			unit_price
		FROM 
			products
		WHERE
			product_id = ?;`
	return query
}

func DeleteProductFromCart() string {
	query :=
		`DELETE FROM users_cart
		WHERE
			product_id = ? AND user_id = ?;`
	return query
}

func IncreaseProductInCart() string {
	query :=
		`UPDATE users_cart
		SET
			count = count + 1,
			total_price = (count + 1) * unit_price
		WHERE
			product_id = ? AND user_id = ?;`
	return query
}

func DecreaseProductInCart() string {
	query :=
		`WITH updatedCnt AS (
			UPDATE users_cart
			SET
				count = count - 1
			WHERE
			product_id = ? AND user_id = ?
		)
		DELETE FROM users_cart
		WHERE	
			count = 0 AND product_id = ? AND user_id = ?;`
	return query
}
