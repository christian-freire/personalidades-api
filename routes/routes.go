package routes

import (
	"api-go-rest/controllers"
	"log"
	"net/http"
)

// Função de iniciação do servidor, que controla as requisições. O HandleFunc especifica
// o endpoint que iremos utilizar na URL e a função que será executada após a chamada do
// endpoint. log.Fatal é equivalente ao Print, mas ele finaliza a execução ao encontrar
// um erro. ListenAndServe indica a porta onde as requisições serão atendidas, no caso :8000
func HandleRequest() {
	http.HandleFunc("/", controllers.Home)
	http.HandleFunc("/api/personalidades", controllers.AllPersonality)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
