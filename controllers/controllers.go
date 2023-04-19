package controllers

import (
	"api-go-rest/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Função da página home, ela chama um "w" que é o ResponseWriter, que irá printar na tela
// nosso comando, e também chama um "r" que é a própria requisição.
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home page")
}

// Função que exebirá os json do array Personality, por meio de json.NewEncoder(w), que irá
// codificar o ResponseWriter para json, e irá codificar a instancia da variável Personalidades,
// assim exibindo os json.
func AllPersonality(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Personalidades)
}

// Função que exibirá apenas uma das personalidades, que terá seu ID especificado na URL
func ReturnOnePersonality(w http.ResponseWriter, r *http.Request) {
	// A função mux.Vars do Gorilla Mux compara as solicitações recebidas com uma lista de
	// rotas registradas e chama um manipulador para a rota que corresponde à URL ou outras condições
	vars := mux.Vars(r)
	id := vars["id"]

	// Criamos um loop for para buscar nos nossos ids qual tem o valor que está sendo buscado na URL,
	// este valor está dentro da variável id acima.
	for _, personalidade := range models.Personalidades {
		if strconv.Itoa(personalidade.Id) == id {
			json.NewEncoder(w).Encode(personalidade)
		}
	}
}
