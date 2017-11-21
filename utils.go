package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/gorilla/mux"
)

func createRoutes(r *mux.Router) {
	r.HandleFunc("/api/urls", urlsHandler).Methods("GET", "POST")
	r.HandleFunc("api/urls/{id}", urlsDetailsHandler).Methods("GET", "DELETE")
}

func createRouter() *mux.Router {
	router := mux.NewRouter()
	createRoutes(router)
	return router
}

func logRequest(r *http.Request) {
	log.Printf("[*] [%s] | [%s] | [%s]", r.Method, r.URL, r.Host)
}

func makeResponse(w http.ResponseWriter, r *http.Request, p interface{}, code int) {
	logRequest(r)
	response, _ := json.Marshal(p)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func makePayload(msg string) map[string]string {
	payload := map[string]string{"message": msg}
	return payload
}

func fetchUrlTitle(url string) string {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}

	title := doc.Find("title").Text()
	return title
}
