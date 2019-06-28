package database

func CreateTable() error {
	stmt := `
	CREATE TABLE IF NOT EXISTS todos(
		id SERIAL PRIMARY KEY,
		title TEXT,
		status TEXT
	);
	`
	_, err := db.Exec(stmt)
	return err
}
