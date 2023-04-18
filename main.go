package main

import (
	"api-go-rest/models"
	"api-go-rest/routes"
	"fmt"
)

// Função main, onde chamamos a função HandleRequest que será iniciada no carregamento da API.
func main() {
	//Instanciando informações ao array Personality,
	//então a variável Personalidades recebe o Array
	//Personality que está instanciado com algunas
	//informações agora.
	models.Personalidades = []models.Personality{
		{Nome: "Nome 1", Historia: "Historia 1"},
		{Nome: "Nome 2", Historia: "Historia 2"},
	}

	fmt.Println("Iniciando API em Go")
	routes.HandleRequest()
}
