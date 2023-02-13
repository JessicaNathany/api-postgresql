package models

import (
	"github.com/aprendagolang/api-pgsql/db"
)

func GetAll() (todos []Todo, err error) {
	conn, err := db.OpenConnection()

	if err != nil {
		return
	}

	defer conn.Close()

	// seleciona todos os registros do banco
	rows := conn.Query(`SELECT * FROM todos`)

	for rows.Next() {

		// percorre a lista de todos os itens retornados
		var todo Todo
		err = rows.Scan(&todo.Id, &todo.Title, &todo.Description, &todo.Done)

		if err != nil {
			continue
		}

		todos = append(todos, todo)
	}
	return
}
