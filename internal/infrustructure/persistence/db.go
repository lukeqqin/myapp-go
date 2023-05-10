package persistence

import (
	"gorm.io/gorm"
)

var GenealogyRepository *_GenealogyRepository

func NewRepositories(db *gorm.DB) {
	GenealogyRepository = NewGenealogyRepository(db)
}
