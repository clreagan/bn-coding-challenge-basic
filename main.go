package main

import (
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		http.ServeFile(res, req, "./index.html")
	})

	http.HandleFunc("/css", func(res http.ResponseWriter, req *http.Request) {
		http.ServeFile(res, req, "./css/style.css")
	})

	http.HandleFunc("/js", func(res http.ResponseWriter, req *http.Request) {
		http.ServeFile(res, req, "./js/script.js")
	})

	log.Fatal(http.ListenAndServe(":8000", nil))
}
