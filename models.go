package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"time"
)

type URL struct {
	ID        int    `gorm:"AUTO_INCREMENT"`
	URL       string `gorm:"primary_key"`
	Title     string
	CreatedAt time.Time
}

func createUrl(db *gorm.DB, url string) {
	db.AutoMigrate(&URL{})
	title := fetchUrlTitle(url)
	db.Create(&URL{URL: url, Title: title})
}

func getUrls(db *gorm.DB) []URL {
	var urls []URL
	db.Find(&urls)
	return urls
}

func getUrl(db *gorm.DB, id int) *URL {
	url := &URL{}
	db.Where(&URL{ID: id}).First(&url)
	return url
}

func deleteUrl(db *gorm.DB, id int) {
	url := &URL{}
	db.Where(&URL{ID: id}).First(&url)
	db.Delete(&url)
}
