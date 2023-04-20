package controllers

import (
	"api-go-rest/database"
	"api-go-rest/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Função da página home, ela chama um "w" que é o ResponseWriter, que irá printar na tela
// nosso comando, e também chama um "r" que é a própria requisição.
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home page")
}

// Função que será responsável por eixibir todas as personalidades cadastradas no banco de
// dados. Criamos uma variável p do tipo models.Personality, e com o comando database.DB.Find,
// procuramos no nosso banco de dados todos os modelos que combinam com a condição, que no caso
// é a própria struct de Personality.
func AllPersonalities(w http.ResponseWriter, r *http.Request) {
	var p []models.Personalidade
	database.DB.Find(&p)
	json.NewEncoder(w).Encode(p)
}

// Função que exibirá apenas uma das personalidades, que terá seu ID especificado na URL
func ReturnOnePersonality(w http.ResponseWriter, r *http.Request) {
	// A função mux.Vars do Gorilla Mux compara as solicitações recebidas com uma lista de
	// rotas registradas e chama um manipulador para a rota que corresponde à URL ou outras condições
	vars := mux.Vars(r)
	id := vars["id"]

	// Agora para retornar apenas uma das personalidades de acordo com nossa URL, criamos uma variável
	// chamada personalidade, que é do tipo models.Personalidade. A função First localiza o primeiro
	// registro ordenado pela chave primária, correspondendo às condições dadas, então ele irá localizar
	// o Id da pessoa, e verá se corresponde ao id da URL.
	var personalidade models.Personalidade
	database.DB.First(&personalidade, id)

	json.NewEncoder(w).Encode(&personalidade)
}
