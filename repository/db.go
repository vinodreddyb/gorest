package repository

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"restexample/configs"
	models2 "restexample/models"
	"time"
)

var DB *gorm.DB

var err error

func Connect() (*gorm.DB, error) {

	host := configs.Config.GetString("DB.ADDRESS")
	port := configs.Config.GetString("DB.PORT")
	user := configs.Config.GetString("DB.USERNAME")
	password := configs.Config.GetString("DB.PASSWORD")
	dbName := configs.Config.GetString("DB.DATABASE")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbName, port)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,
		},
	)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	//Migrate repository structure into the db
	err = DB.AutoMigrate(&models2.User{})
	if err != nil {
		return nil, err
	}

	return DB, nil
}
