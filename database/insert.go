package database

func InsertTodo(title string, status string) (int, error) {
	query := `
	INSERT INTO todos (title, status) VALUES ($1, $2) RETURNING id;
	`
	var id int
	row := db.QueryRow(query, title, status)
	err := row.Scan(&id)
	return id, err
}
