package main

import (
	"net/http"
	// Importação do pacote handlers local (como deve ser no Go)
	"software-finance/backend/database"
	"software-finance/backend/handlers" 

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// A lógica de godotenv.Load() foi removida pois as variáveis de ambiente 
	// são fornecidas diretamente pelo docker-compose.yml.

	// 1. Conecta com o Banco de Dados
	database.Conectar()

	router := gin.Default()

	// 2. Configuração de CORS (Essencial para React se comunicar com Go)
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true 
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept"}
	router.Use(cors.New(config))

	// Rota de Teste Simples (mantida)
	router.GET("/api/status", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Finance API Running in Go!"})
	})
	
	// 3. Novas Rotas da API para Transações
	// POST /api/transacoes/criar: Rota de criação (para evitar conflito com o GET /api/transacoes)
	router.POST("/api/transacoes/criar", handlers.CriarTransacao) 

	// GET /api/transacoes: Rota de listagem
	router.GET("/api/transacoes", handlers.ListarTransacoes) 

	// Roda a API na porta 8080 (Será mapeada pelo Docker)
	router.Run(":8080")
}






