package controller

import (
	"fmt"
	"net/http"
	"project-for-portfolioDEV/GOAPI/src/models"
	"project-for-portfolioDEV/GOAPI/src/service"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type NovalController struct {
	novalService service.INovalService
}

func NewNovalController(novalService service.INovalService) *NovalController {
	return &NovalController{
		novalService: novalService,
	}
}

func (controller *NovalController) CraeteNoval(ctx *fiber.Ctx) error {
	var (
		novalReqesut  models.Noval
		novalResponse models.Response
	)

	requestNoval := new(models.Noval)
	//handle the request
	// Parse the request body into the Noval struct
	if err := ctx.BodyParser(&requestNoval); err != nil {
		novalResponse = models.Response{
			StatusCode: http.StatusBadRequest,
			Message:    fmt.Sprintf("Failed to Parse Request Body with %s", err.Error()),
		}
		return ctx.Status(http.StatusBadRequest).JSON(novalResponse)
	}

	//validate the struct fields
	validate := validator.New()
	if err := validate.Struct(requestNoval); err != nil {
		novalResponse = models.Response{
			StatusCode: http.StatusBadRequest,
			Message:    fmt.Sprintf("Failed to Validate Struct Body with %s", err.Error()),
		}
		return ctx.Status(http.StatusBadRequest).JSON(novalResponse)

	}

	//save to database
	if err := controller.novalService.CreateNoval(novalReqesut); err != nil {
		novalResponse = models.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
		return ctx.Status(http.StatusInternalServerError).JSON(novalResponse)
	}

	return ctx.JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": "Noval successfully created",
	})
}
