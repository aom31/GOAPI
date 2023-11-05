package controller

import (
	"fmt"
	"net/http"
	"project-for-portfolioDEV/GOAPI/src/domain"
	"project-for-portfolioDEV/GOAPI/src/models"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type DocumentController struct {
	documentService domain.DocumentService
}

func NewDocumentController(documentService domain.DocumentService) *DocumentController {
	return &DocumentController{
		documentService: documentService,
	}
}

func (controller *DocumentController) CreateDocument(ctx *fiber.Ctx) error {
	var (
		docRequest models.Document
		Response   models.Response
	)

	//handle the request
	// Parse the request body into the Noval struct
	if err := ctx.BodyParser(&docRequest); err != nil {
		Response = models.Response{
			StatusCode: http.StatusBadRequest,
			Message:    fmt.Sprintf("Failed to Parse Request Body with %s", err.Error()),
		}
		return ctx.Status(http.StatusBadRequest).JSON(Response)
	}

	//validate the struct fields
	validate := validator.New()
	if err := validate.Struct(docRequest); err != nil {
		Response = models.Response{
			StatusCode: http.StatusBadRequest,
			Message:    fmt.Sprintf("Failed to Validate Struct Body with %s", err.Error()),
		}
		return ctx.Status(http.StatusBadRequest).JSON(Response)

	}

	//save to database
	if err := controller.documentService.CreateDocument(docRequest); err != nil {
		Response = models.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
		return ctx.Status(http.StatusInternalServerError).JSON(Response)
	}

	return ctx.JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": "Noval successfully created",
	})
}
