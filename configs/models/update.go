package models

import (
	"github.com/aprendagolang/api-pgsql/db"
)

func Update(id int64, todo Todo) (int64, error) {
	conn, err := db.OpenConnection()

	if err != nil {
		return 0, err
	}

	defer conn.Close()

	res, err := conn.Exec(`UPDATE todos SET title=$2, description=$4, done=$3, WHERE id=$1`, id, todo.Title, todo.Description, todo.Done, todo.Id)

	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
