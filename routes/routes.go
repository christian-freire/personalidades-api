package routes

import (
	"api-go-rest/controllers"
	"api-go-rest/middleware"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// Função de iniciação do servidor, que controla as requisições. O HandleFunc especifica
// o endpoint que iremos utilizar na URL e a função que será executada após a chamada do
// endpoint. log.Fatal é equivalente ao Print, mas ele finaliza a execução ao encontrar
// um erro. ListenAndServe indica a porta onde as requisições serão atendidas, no caso :8000
func HandleRequest() {
	// Usando o framework Gorilla Mux para a criação do Router
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidades", controllers.AllPersonalities).Methods("Get")
	r.HandleFunc("/api/personalidades/{id}", controllers.ReturnOnePersonality).Methods("Get")
	r.HandleFunc("/api/personalidades", controllers.CreateNewPersonality).Methods("Post")
	r.HandleFunc("/api/personalidades/{id}", controllers.DeletePersonality).Methods("Delete")
	r.HandleFunc("/api/personalidades/{id}", controllers.EditPersonality).Methods("Put")
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
