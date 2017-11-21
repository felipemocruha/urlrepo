package main

import (
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func main() {
	router := createRouter()

	const addr = "postgresql://maxroach@localhost:26257/urlrepo?sslmode=disable"
	db, err := gorm.Open("postgres", addr)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	log.Println("[*] Server started...")
	log.Fatal(http.ListenAndServe(":8000", router))
}
