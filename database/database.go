package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"

	utils "bitbucket.org/envirovisionsolutions/showandtell/utils"
)

type Db struct {
	Db *gorm.DB
}

var (
	DB         Db
	baseFolder = "./storage/db"
)

func ConnectDb() {
	appEnv := utils.EnvGet("APP_ENV", "production")
	logLevel := logger.Error
	if appEnv == "development" {
		logLevel = logger.Info
	}
	log.Printf("INFO: DB Log level: %d\n", logLevel)
	if _, err := os.Stat(baseFolder); os.IsNotExist(err) {
		os.MkdirAll(baseFolder, 0777)
	}
	db, err := gorm.Open(sqlite.Open("storage/db/showandtell.sqlite"), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	}
	db.Logger = logger.Default.LogMode(logLevel)

	DB = Db{Db: db}
}
