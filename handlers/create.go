package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func Create(httpResponse http.ResponseWriter, request *http.Request) {
	var todo models.Todo

	err := json.NewDecoder(request.Body).Decode(&todo)

	if err != nil {
		log.Println("Erro ao fazer decode do json")
		http.Error(httpResponse, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	id, err := models.Insert(todo)

	var resp map[string]any

	if err != nil{
		resp = map[string]any{
			"Error" : true,
			"Message": fmt.Sprintf("Ocorreu um erro ao tentar inserir: %v", err),
		}
	} else{
		resp = map[string]any{
			"Error": false,
			"Message": fmt.Sprintf("Todo inserido com sucesso ID:! %d", id),
		}
		}
	}

}
