package query

func CheckPermission() string {
	query :=
		`SELECT EXISTS (
			SELECT 1
			FROM 
				product_rules pr
			JOIN 
				users_groups ON pr.group_id = ug.group_id
			WHERE
				ug.user_id = ? AND pr.rule = ?
		);`
	return query
}
