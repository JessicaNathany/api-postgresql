package models

import (
	"github.com/aprendagolang/api-pgsql/db"
)

func Delete(id int64) (int64, error) {
	conn, err := db.OpenConnection()

	if err != nil {
		return 0, err
	}

	defer conn.Close()

	res, err := conn.Exec(`DELETE FROM todos WHERE id=$1`, id, todo.Title, todo.Description, todo.Done, todo.Id)

	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
