package database

func DeleteTodo(id int) error {
	stmt, err := db.Prepare("DELETE FROM todos WHERE id=$1;")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(id)
	return err
}
