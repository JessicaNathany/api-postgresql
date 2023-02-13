package db

import (
	"crypto/aes"
	"database/sql"
	"fmt"
	"github.com/aprendagolang/api-psgql/configs"
	_ "github.com/lib/pq"
)

// abre conexão com o banco de dados
func OpenConnection() (*sql.DB, error) {
	config := configs.GetDB()

	stringConnection := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable" 
	conf.Host, conf.Port, conf.User, conf.Password, conf.Database)

	connection, erro := sql.Open("postgress", stringConnection)

	if erro != nil {
		panic(erro)
	}

	// faz um ping no banco pra ver se a conexão foi estabelecida
	erro = connection.Ping()

	return connection, erro
}

