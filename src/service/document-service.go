package service

import (
	"project-for-portfolioDEV/GOAPI/src/domain"
	"project-for-portfolioDEV/GOAPI/src/models"
)

type documentService struct {
	docRepo domain.DocumentRepo
}

func NewDocumentService(documentRepo domain.DocumentRepo) domain.DocumentService {
	return &documentService{
		docRepo: documentRepo,
	}
}

func (service *documentService) CreateDocument(doc models.Document) error {
	if err := service.docRepo.CreateDocument(doc); err != nil {
		return err
	}

	return nil
}
