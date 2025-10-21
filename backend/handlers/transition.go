package handlers

import (
	"net/http"
	"software-finance/backend/database"
	"software-finance/backend/models"

	"github.com/gin-gonic/gin"
)

// ListarTransacoes retorna todas as transações do banco de dados.
func ListarTransacoes(c *gin.Context) {
	var transacoes []models.Transacao
	// Busca todas as transações no DB, ordenando pela data de criação
	result := database.DB.Order("created_at desc").Find(&transacoes)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar transações"})
		return
	}

	c.JSON(http.StatusOK, transacoes)
}

// CriarTransacao recebe dados do Frontend e insere uma nova transação no DB.
func CriarTransacao(c *gin.Context) {
	var novaTransacao models.Transacao

	// Bind JSON: Tenta mapear o JSON recebido para a struct models.Transacao
	if err := c.ShouldBindJSON(&novaTransacao); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados de transação inválidos"})
		return
	}

	// Validação básica
	if novaTransacao.Valor == 0 || novaTransacao.Descricao == "" || (novaTransacao.Tipo != "Receita" && novaTransacao.Tipo != "Despesa") {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Campos obrigatórios faltando ou inválidos (Valor, Descrição, Tipo)"})
		return
	}

	// Insere no banco de dados
	result := database.DB.Create(&novaTransacao)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao salvar transação no banco de dados"})
		return
	}

	// Retorna a transação criada (incluindo o ID gerado pelo DB)
	c.JSON(http.StatusCreated, novaTransacao)
}
