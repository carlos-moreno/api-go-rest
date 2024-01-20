package main

import (
	"fmt"

	"github.com/carlos-moreno/api-go-rest/models"
	"github.com/carlos-moreno/api-go-rest/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Nome: "Nome 1", Historia: "História 1"},
		{Nome: "Nome 2", Historia: "História 2"},
	}

	fmt.Println("Iniciando o servidor Rest com Go")
	routes.HandleRequest()
}
