package database

import (
	"fmt"
	"log"
	"os"

	"software-finance/backend/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Conectar() {
	// Pega variáveis de ambiente definidas no docker-compose.yml
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", 
		host, user, password, dbname, port)

	var err error
	// Tenta conectar com o banco de dados
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		// Log e encerra o programa se a conexão falhar
		log.Fatalf("Falha ao conectar com o banco de dados: %v", err)
	}

	fmt.Println("Conexão com o banco de dados estabelecida com sucesso!")

	// Migra o esquema (cria a tabela 'transacoes' se não existir)
	err = DB.AutoMigrate(&models.Transacao{})
	if err != nil {
		log.Fatalf("Falha ao migrar o esquema: %v", err)
	}
	fmt.Println("Migração do esquema concluída com sucesso!")
}