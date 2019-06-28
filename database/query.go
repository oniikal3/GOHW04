package database

import "database/sql"

func QueryTodo(id int) (*sql.Row, error) {
	stmt, err := db.Prepare("SELECT id, title, status FROM todos WHERE id=$1;")
	row := stmt.QueryRow(id)
	return row, err
}

// func (db *sql.DB) QueryTodo(id int) (*sql.Row, error) {
// 	stmt, err := db.Prepare("SELECT id, title, status FROM todos WHERE id=$1;")
// 	row := stmt.QueryRow(id)
// 	return row, err
// }
