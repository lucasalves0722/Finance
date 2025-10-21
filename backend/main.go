package main

import (
	"net/http"
	"software-finance/backend/database"
	"software-finance/backend/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// 1. Conecta com o Banco de Dados
	database.Conectar()

	router := gin.Default()

	// 2. Configuração de CORS (Essencial para React se comunicar com Go)
	// Permite qualquer origem e método para desenvolvimento
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true 
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept"}
	router.Use(cors.New(config))

	// Rota de Teste Simples (mantida)
	router.GET("/api/status", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "API de Finanças Rodando em Go!"})
	})
	
	// 3. Novas Rotas da API para Transações
	router.POST("/api/transacoes", handlers.CriarTransacao) // Inserir nova
	router.GET("/api/transacoes", handlers.ListarTransacoes) // Listar todas

	// Roda a API na porta 8080 (Será mapeada pelo Docker)
	router.Run(":8080")
}