package router

import (
	"log"
	"project-for-portfolioDEV/GOAPI/src/controller"

	"github.com/gofiber/fiber/v2"
)

func NewRouter(router *fiber.App, documentController *controller.DocumentController) *fiber.App {

	router.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("hello word")
	})
	router.Post("/document", documentController.CreateDocument)
	log.Println("starting route POST: /document")

	return router

}
