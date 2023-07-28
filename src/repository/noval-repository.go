package repository

import (
	"errors"
	"project-for-portfolioDEV/GOAPI/src/models"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type INovalRepository interface {
	CreateNoval(createNoval models.Noval) error
}

type novalRepository struct {
	db    *gorm.DB
	redis *redis.Client
}

func NewNovalRepository(db *gorm.DB, redis *redis.Client) INovalRepository {
	return &novalRepository{
		db:    db,
		redis: redis,
	}
}

func (repo *novalRepository) CreateNoval(createNoval models.Noval) error {

	if err := repo.db.Create(&createNoval).Error; err != nil {
		return errors.New("failed to create noval")
	}

	return nil
}
