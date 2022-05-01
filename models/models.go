package models

import (
	"fmt"
	"log"
	"yslinear/go-covid19/pkg/setting"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

// Setup initializes the database instance
func Setup() {
	var err error

	db, err = gorm.Open(postgres.Open(dsn("postgres")), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	if err := db.Exec("CREATE DATABASE " + setting.DatabaseSetting.Name); err != nil {
		log.Printf("[info] database %s created", setting.DatabaseSetting.Name)
		db, _ = gorm.Open(postgres.Open(dsn(setting.DatabaseSetting.Name)), &gorm.Config{})
		db.AutoMigrate(&Hospital{}, &Fst{})
	} else {
		log.Printf("[info] database %s already exist", setting.DatabaseSetting.Name)
		if err != nil {
			panic(err)
		}
	}

	db, err = gorm.Open(postgres.Open(dsn(setting.DatabaseSetting.Name)), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()

	if err != nil {
		log.Fatalf("sqlDB err: %v", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
}

func dsn(name string) string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s",
		setting.DatabaseSetting.Host,
		setting.DatabaseSetting.User,
		setting.DatabaseSetting.Password,
		name,
		setting.DatabaseSetting.Port,
		setting.DatabaseSetting.TimeZone,
	)
}

// CloseDB closes database connection (unnecessary)
func CloseDB() {
	sqlDB, err := db.DB()

	if err != nil {
		log.Fatalf("sqlDB err: %v", err)
	}
	defer sqlDB.Close()
}
