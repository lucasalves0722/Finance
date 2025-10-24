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
	// Pega variáveis de ambiente
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	// Garante que o port não está vazio, para evitar o erro de parsing
    if port == "" {
        port = "5432" // Valor de fallback, caso esteja vazio (ajuste se necessário)
        fmt.Printf("Attention: DB_PORT variable not defined. Using default port: %s\n", port)
    }

	// Monta a string de conexão (DSN)
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", 
		host, user, password, dbname, port)

	var err error
	// Tenta conectar com o banco de dados
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		// Log e encerra o programa se a conexão falhar
		log.Fatalf("Failed to connect to the database: %v. Check your environment variables: %s", err, dsn)
	}

	fmt.Println("Connection to the database established successfully!")

	// Migra o esquema (cria a tabela 'transacoes' se não existir)
	err = DB.AutoMigrate(&models.Transacao{})
	if err != nil {
		log.Fatalf("Failed to migrate the schema: %v", err)
	}
	fmt.Println("Schema migration completed successfully!")
}