package handlers

import (
	"fmt"
	"net/http"
	"os"
	"software-finance/backend/database"
	"software-finance/backend/models"

	"github.com/gin-gonic/gin"
)

// Struct auxiliar para receber os dados do Frontend.
// O campo Valor agora é float64, pois o Frontend está enviando um número JSON válido.
type TransacaoInput struct {
	Descricao string  `json:"descricao" binding:"required"`
	Valor     float64 `json:"valor" binding:"required"` // Corrigido para float64
	Tipo      string  `json:"tipo" binding:"required"`
}

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
	var input TransacaoInput
	
	// 1. Bind para a struct auxiliar (recebe o Valor como float64)
	if err := c.ShouldBindJSON(&input); err != nil {
		// Se o Bind falhar (ex: por causa de um tipo inesperado), logamos a falha.
		fmt.Fprintf(os.Stderr, "Erro de Bind no JSON: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados de transação inválidos. Verifique se o Valor é um número."})
		return
	}

	// 2. Validação final dos dados
	if input.Valor <= 0 || input.Descricao == "" || (input.Tipo != "Receita" && input.Tipo != "Despesa") {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Campos obrigatórios faltando ou inválidos (Valor deve ser > 0, Descrição, Tipo)."})
		return
	}

	// 3. Cria o objeto final para o DB
	novaTransacao := models.Transacao{
		Descricao: input.Descricao,
		Valor:     input.Valor, // Agora é diretamente o float64 recebido
		Tipo:      input.Tipo,
	}

	// 4. Insere no banco de dados
	result := database.DB.Create(&novaTransacao)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao salvar transação no banco de dados."})
		return
	}

	// Retorna sucesso
	c.JSON(http.StatusCreated, novaTransacao)
}
