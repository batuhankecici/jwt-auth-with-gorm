package managers

import (
	"jwt-auth-with-gorm/models"

	"github.com/jinzhu/gorm"
)

type BookStorage interface {
	Get(id int) (*models.Book, error)
	//Create(*models.Book) error
}

type BookManager struct {
	db *gorm.DB
}

func NewBookManager(conn *gorm.DB) BookStorage {
	return &BookManager{db: conn}
}

func (b BookManager) Get(id int) (*models.Book, error) {
	books := &models.Book{}
	if err := b.db.First(&books, id).Error; err != nil {
		return nil, err
	}
	return books, nil
}
