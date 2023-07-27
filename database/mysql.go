package database

import (
	"fmt"
	"log"
	"project-for-portfolioDEV/GOAPI/config"

	"gorm.io/driver/mysql"

	"gorm.io/gorm"
)

func ConnectionMySQLDB(config *config.Config) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.DBUsername, config.DBPassword, config.DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil
	}

	log.Println("Success Connected to database Mysql with", config.DBName)

	return db
}
