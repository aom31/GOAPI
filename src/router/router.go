package router

import (
	"project-for-portfolioDEV/GOAPI/src/controller"

	"github.com/gofiber/fiber/v2"
)

func NewRouter(router *fiber.App, novalController *controller.NovalController) *fiber.App {

	router.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("hello word")
	})
	router.Post("/noval", novalController.CraeteNoval)

	return router

}
