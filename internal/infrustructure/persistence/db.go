package persistence

import (
	"gorm.io/gorm"
)

var GenealogyRepository *_GenealogyRepository
var UserRepository *_UserRepository
var GenealogyTagRepository *_GenealogyTagRepository

func NewRepositories(db *gorm.DB) {
	GenealogyRepository = newGenealogyRepository(db)
	UserRepository = newUserRepository(db)
	GenealogyTagRepository = newGenealogyTagRepository(db)
}
