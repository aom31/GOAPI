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
	dbMysql.AutoMigrate(&models.Noval{})

	//redis
	dbRedis := database.ConnectionRedisDb(&loadConfig)

	startServer(dbMysql, dbRedis)
}

func startServer(dbMysql *gorm.DB, dbRedis *redis.Client) {
	app := fiber.New()

	novalRepo := repository.NewNovalRepository(dbMysql, dbRedis)
	novalService := service.NewNovalService(novalRepo)
	novalController := controller.NewNovalController(novalService)

	routers := router.NewRouter(app, novalController)

	err := routers.Listen(":3400")
	if err != nil {
		panic(err)
	}

}
