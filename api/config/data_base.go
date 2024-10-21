package conf

import (
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"koriebruh/restful/api/model/domain"
	"log"
	"os"
	"time"
)

func InitDB() *gorm.DB {
	//#MAKE CONNECTION
	//
	err := godotenv.Load()
	if err != nil {
		panic("Failed to load .env file")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSLMODE"),
		os.Getenv("DB_TIMEZONE"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(errors.New("Failed Connected into data base"))
	}
	log.Println("success connect database")

	//#CONNECTION POOL
	//
	sqlDB, err := db.DB()
	if err != nil {
		panic(errors.New("Failed to configure database connection pool"))
	}
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetConnMaxLifetime(time.Hour)
	sqlDB.SetConnMaxIdleTime(time.Minute * 5)
	log.Println("success setup connection pool")

	//#AUTOMIGRATE
	//
	if err = db.AutoMigrate(
		&domain.User{},
	); err != nil {
		panic(errors.New("Failed to migrate model"))
	}
	log.Println("success auto migrate database")

	return db

}
