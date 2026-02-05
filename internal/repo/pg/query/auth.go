package query

func GetUserByLogPass() string {
	query :=
		`SELECT
			u.user_id
		FROM 
			users u
		WHERE
			u.user_login = ? AND u.user_password = ?`
	return query
}
