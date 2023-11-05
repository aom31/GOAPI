package repository

import (
	"errors"
	"project-for-portfolioDEV/GOAPI/src/domain"
	"project-for-portfolioDEV/GOAPI/src/models"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type documentRepository struct {
	db    *gorm.DB
	redis *redis.Client
}

func NewDocumentRepository(db *gorm.DB, redis *redis.Client) domain.DocumentRepo {
	return &documentRepository{
		db:    db,
		redis: redis,
	}
}

func (repo *documentRepository) CreateDocument(doc models.Document) error {

	if err := repo.db.Create(&doc).Error; err != nil {
		return errors.New("failed to create document")
	}

	return nil
}
