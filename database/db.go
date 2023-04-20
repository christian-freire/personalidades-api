package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Variável que inicia o gorm no projeto, que vai ser responsável por conectar nossa API ao banco
// de dados postgres
var (
	DB  *gorm.DB
	err error
)

// Função nos conecta ao banco de dados com as chaves que estão na variável stringDeConexao, que foram
// definidas na nossa imagem postgres que está no conteiner Docker, no arquivo docker-compose.yml
func ConectWithDataBase() {
	stringDeConexao := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
	// Aqui iniciamos a conexão com o DB com a funação gorm.Open
	DB, err = gorm.Open(postgres.Open(stringDeConexao))
	if err != nil {
		log.Panic("Erro ao conectar com banco de dados")
	}
}
