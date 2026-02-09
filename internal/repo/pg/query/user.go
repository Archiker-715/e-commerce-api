package query

func GetUserByLogPass() string {
	query :=
		`SELECT
			u.user_id
		FROM 
			users u
		WHERE
			u.user_login = ? AND u.user_password = ?;`
	return query
}

func CreateUser() string {
	query :=
		`INSERT INTO users(
			user_id,
			user_login,
			user_password,
			inserted_by,
			inserted
		)
		SELECT 
			?,
			?,
			?,
			?,
			?;`
	return query
}
