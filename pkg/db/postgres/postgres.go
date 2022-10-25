package postgres

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"rest-gin/internal/video"
	"time"
)

var dbase *gorm.DB

func Init() *gorm.DB {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  "host=localhost user=admin password=1254 dbname=postgres port=5454 sslmode=disable",
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&video.Video{})
	return db
}

func GetDB() *gorm.DB {
	if dbase == nil {
		dbase = Init()
		var sleep = time.Duration(1)
		for dbase == nil {
			sleep = sleep * 2
			fmt.Printf("Database is unavailable. Wait fof %d second", sleep)
			time.Sleep(sleep * time.Second)
			dbase = Init()
		}

	}
	return dbase
}
