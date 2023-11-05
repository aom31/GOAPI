package domain

import "project-for-portfolioDEV/GOAPI/src/models"

type DocumentRepo interface {
	CreateDocument(doc models.Document) error
}

type DocumentService interface {
	CreateDocument(doc models.Document) error
}
