package main

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"os"
)

func fetchUrlTitle(url string) string {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Println(err)
		return ""
	}

	title := doc.Find("title").Text()
	return title
}

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}
