package persistence

import (
	"gorm.io/gorm"
)

var GenealogyRepository *_GenealogyRepository
var UserRepository *_UserRepository
var GenealogyMembersRepository *_GenealogyMembersRepository
var CosRepository *_CosRepository

func NewRepositories(db *gorm.DB) {
	GenealogyRepository = newGenealogyRepository(db)
	UserRepository = newUserRepository(db)
	GenealogyMembersRepository = newGenealogyMembersRepository(db)
	CosRepository = NewCosRepository(db)
}
