package controllers

import (
	"api-go-rest/models"
	"encoding/json"
	"fmt"
	"net/http"
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
