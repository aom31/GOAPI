package main

import (
	"log"
	"project-for-portfolioDEV/GOAPI/config"
	"project-for-portfolioDEV/GOAPI/database"

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
	dbRedis := database.ConnectionRedisDb(&loadConfig)

	startServer(dbMysql, dbRedis)
}

func startServer(dbMysql *gorm.DB, dbRedis *redis.Client) {
	app := fiber.New()

	err := app.Listen(":3400")
	if err != nil {
		panic(err)
	}

}
