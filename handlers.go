package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func urlsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		urls := getUrls(db)
		makeResponse(w, r, urls, 200)

	} else if r.Method == "POST" {
		payload := makePayload("created")
		makeResponse(w, r, payload, 201)
	}
}

func urlsDetailsHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	log.Println(id)
	if err != nil {
		log.Fatal(err)
	}

	if r.Method == "GET" {
		url := getUrl(db, id)
		makeResponse(w, r, url, 200)

	} else if r.Method == "DELETE" {
		payload := makePayload("")
		makeResponse(w, r, payload, 204)
	}
}
