package main

import (
	"fmt"
	"log"
	"net/http"
)

// Função da página home, ela chama um "w" que é o ResponseWriter, que irá printar na tela
// nosso comando, e também chama um "r" que é a própria requisição.
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home page")
}

// Função de iniciação do servidor, que controla as requisições. HandleFunc indica o caminho
// na URL que queremos e executa a função "Home" no caso. Logo após temos um log.Fatal, que
// indica que após qualquer erro encontrado o processo deve ser encerrado, e inicializamos
// o servidor na porta :8000
func HandleRequest() {
	http.HandleFunc("/", Home)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

// Função main, onde chamamos a função HandleRequest que será iniciada no carregamento da API.
func main() {
	fmt.Println("Hello World")
	HandleRequest()
}
