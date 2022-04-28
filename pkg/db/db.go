package db

import (
	"log"
	"os"
	"yslinear/go-covid19/pkg/setting"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Setup() {
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

	_, err := gorm.Open(sqlite.Open(GetSQliteSavePath()), &gorm.Config{})
	if err != nil {
		log.Fatalf("gorm.Open err: %v", err)
	}
}

func GetSQliteSavePath() string {
	return setting.DatabaseSetting.SqlitePath
}
