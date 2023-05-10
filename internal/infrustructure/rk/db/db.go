package db

import (
	"gorm.io/gorm"
	"myapp-go/internal/infrustructure/persistence"
)

func New(db *gorm.DB) {
	persistence.NewRepositories(db)
}
