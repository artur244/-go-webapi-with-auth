package database

import (
	"log"
	"time"

	"github.com/artur244/first-go-rest-api/database/migrations"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func StartDB() {
	// access := "host=localhost port=25432 user=admin dbname=books sslmode=disable password=123456"
	// database, error := gorm.Open(postgres.Open(access), &gorm.Config{})
	access := "root:@tcp(127.0.0.1:3306)/books?charset=utf8mb4&parseTime=True&loc=Local"
	database, error := gorm.Open(mysql.Open(access), &gorm.Config{})

	if error != nil {
		log.Fatal("error: ", error)
	}

	db = database

	config, _ := db.DB()

	config.SetMaxIdleConns(10)
	config.SetMaxOpenConns(100)
	config.SetConnMaxLifetime(time.Hour)

	migrations.RunMigrations(db)
}

func GetDatabase() *gorm.DB {
	return db
}
