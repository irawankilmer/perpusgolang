package config

import (
	"fmt"
	"github.com/irawankilmer/perpustakaanonline/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

var DB *gorm.DB

func InitDB() {
	var err error
	dbport := os.Getenv("DB_PORT")
	dbhost := os.Getenv("DB_HOST")
	dbuser := os.Getenv("DB_USER")
	dbpass := os.Getenv("DB_PASS")
	dbname := os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", dbuser, dbpass, dbhost, dbport, dbname)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Could connect to the database: %v", err)
	}

	log.Println("Database connection established")

	err = DB.AutoMigrate(
		&models.User{},
		&models.Profile{},
	)

	if err != nil {
		log.Fatalf("Could not migrate database: %v", err)
	}
}
