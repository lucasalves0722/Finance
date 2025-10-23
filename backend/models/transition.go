package models

import (
	"gorm.io/gorm"
)

// Transacao representa um item financeiro no banco de dados.
type Transacao struct {
	gorm.Model // Campos automáticos: ID, CreatedAt, UpdatedAt, DeletedAt

	// Adicionado tags JSON para garantir que o nome da chave enviada ao Frontend seja CamelCase (Descricao, Tipo, Valor).
	// Isso resolve o problema de colunas vazias.
	Descricao string `json:"Descricao" gorm:"type:varchar(255);not null"`
	Valor     float64 `json:"Valor" gorm:"type:numeric;not null"`
	Tipo      string `json:"Tipo" gorm:"type:varchar(50);not null"` // Ex: "Receita" ou "Despesa"
}