package models

import (
	"github.com/aprendagolang/api-pgsql/db"
)

func Insert(todo Todo) (id int64, err error) {

	conn, error := db.OpenConnection()

	if err != nil {
		return
	}

	defer conn.Close()

	sql := `INSERT INTO todos (title, description, done) VALUES ($1, $2, $3) RETURNING  id`

	err = conn.QueryRow(sql, todo.Title, todo.Description, todo.Description).Scan(&id)

	return
}
