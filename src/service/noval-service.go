package service

import (
	"project-for-portfolioDEV/GOAPI/src/models"
	"project-for-portfolioDEV/GOAPI/src/repository"
)

type INovalService interface {
	CreateNoval(createNoval models.Noval) error
}

type novalService struct {
	repoNoval repository.INovalRepository
}

func NewNovalService(novalRepo repository.INovalRepository) INovalService {
	return &novalService{
		repoNoval: novalRepo,
	}
}

func (service *novalService) CreateNoval(createNoval models.Noval) error {
	err := service.repoNoval.CreateNoval(createNoval)
	if err != nil {
		return err
	}

	return nil
}
