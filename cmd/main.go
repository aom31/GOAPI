package main

import (
	"log"
	"project-for-portfolioDEV/GOAPI/config"
	"project-for-portfolioDEV/GOAPI/database"
	"project-for-portfolioDEV/GOAPI/src/controller"
	"project-for-portfolioDEV/GOAPI/src/models"
	"project-for-portfolioDEV/GOAPI/src/repository"
	"project-for-portfolioDEV/GOAPI/src/router"
	"project-for-portfolioDEV/GOAPI/src/service"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func main() {

	//connect database and redis
	loadConfig, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load env", err)
	}

	//mysql
	dbMysql := database.ConnectionMySQLDB(&loadConfig)
	dbMysql.AutoMigrate(&models.Document{})

	//redis
	dbRedis := database.ConnectionRedisDb(&loadConfig)

	startServer(dbMysql, dbRedis)
}

func startServer(dbMysql *gorm.DB, dbRedis *redis.Client) {
	app := fiber.New()

	docRepo := repository.NewDocumentRepository(dbMysql, dbRedis)
	docService := service.NewDocumentService(docRepo)
	docController := controller.NewDocumentController(docService)

	routers := router.NewRouter(app, docController)

	err := routers.Listen(":3400")
	if err != nil {
		panic(err)
	}

}
