package models

import (
	"log"
	"os"
	"yslinear/go-covid19/pkg/setting"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

type Model struct {
	ID         int `gorm:"primary_key" json:"id"`
	CreatedOn  int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
	DeletedOn  int `json:"deleted_on"`
}

// Setup initializes the database instance
func Setup() {
	var err error

	if _, err := os.Stat(GetSQliteSavePath()); os.IsNotExist(err) {
		r, err := os.Open("upload/example.db.sqlite")
		if err != nil {
			panic(err)
		}
		defer r.Close()
		w, err := os.Create(GetSQliteSavePath())
		if err != nil {
			panic(err)
		}
		defer w.Close()
		w.ReadFrom(r)
	}

	db, err = gorm.Open(sqlite.Open(GetSQliteSavePath()), &gorm.Config{})

	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}

	sqlDB, err := db.DB()

	if err != nil {
		log.Fatalf("sqlDB err: %v", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
}

// CloseDB closes database connection (unnecessary)
func CloseDB() {
	sqlDB, err := db.DB()

	if err != nil {
		log.Fatalf("sqlDB err: %v", err)
	}
	defer sqlDB.Close()
}

func GetSQliteSavePath() string {
	return setting.DatabaseSetting.SqlitePath
}
