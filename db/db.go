package db

import (
	"../entity"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"os"
)

var (
	db  *gorm.DB
	err error
)

func Init() {
	env := os.Getenv("ENV")

	if "production" == env {
		env = "production"
	} else {
		env = "development"
	}

	// Get Environment Variable
	godotenv.Load(".env." + env)
	godotenv.Load()

	// connect DB
	db, err = gorm.Open("mysql", os.Getenv("CONNECT"))

	if err != nil {
		panic(err)
	}

	autoMigrattion()
}

// Get DB
func GetDB() *gorm.DB {
	return db
}

// force quit
func Close() {
	if err := db.Close(); err != nil {
		panic(err)
	}
}

// migration
func autoMigrattion() {
	db.AutoMigrate(&entity.Product{})
}
