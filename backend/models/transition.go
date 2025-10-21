package models

import (
	"gorm.io/gorm"
)

// Transacao representa um item financeiro no banco de dados.
type Transacao struct {
	gorm.Model // Campos automáticos: ID, CreatedAt, UpdatedAt, DeletedAt
	Descricao string  `json:"descricao" gorm:"type:varchar(255);not null"`
	Valor     float64 `json:"valor" gorm:"type:numeric;not null"`
	Tipo      string  `json:"tipo" gorm:"type:varchar(50);not null"` // Ex: "Receita" ou "Despesa"
}