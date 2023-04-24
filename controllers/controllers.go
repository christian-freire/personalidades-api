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
	// A função w.Header().Set("Content-type", "application/json") indica que a resposta que
	// esperamos no cabeçalho da requisição é do tipo json.
	w.Header().Set("Content-type", "application/json")

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

// Função que cria nova personalidade. Primeiro iniciamos uma variável do tipo models.Personalidade,
// depois decodificamos a requisição que chegou no corpo e instânciamos para dentro da variável
// novaPersonalidade, apontando para ela. Depois criamos a nova personalidade no banco de dados usando
// o database.DB.Create e apontamos para a variável como referência. Por último encodamos a variável
// novaPersonalidade, que está com a nova personalidade, como a resposta que deve chegar na requisição.
func CreateNewPersonality(w http.ResponseWriter, r *http.Request) {
	var novaPersonalidade models.Personalidade
	json.NewDecoder(r.Body).Decode(&novaPersonalidade)
	database.DB.Create(&novaPersonalidade)
	json.NewEncoder(w).Encode(novaPersonalidade)
}

// Basicamente a mesma função do ReturnOnePersonality, a única diferença é que dessa vez utilizaremos
// database.DB.Delete para deletar uma das nossas personalidades.
func DeletePersonality(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var personalidade models.Personalidade
	database.DB.Delete(&personalidade, id)
	json.NewEncoder(w).Encode(personalidade)
}

// Basicamente a mesma função do CreateNewPersonality, a única diferença é que dessa vez utilizaremos
// database.DB.Save para editar uma das nossas personalidades.
func EditPersonality(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var personalidade models.Personalidade
	database.DB.First(&personalidade, id)
	json.NewDecoder(r.Body).Decode(&personalidade)
	database.DB.Save(&personalidade)
	json.NewEncoder(w).Encode(personalidade)
}
