package main

import (
	"golangweb/handler"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handler.HomeHandle)
	mux.HandleFunc("/helo", handler.HeloHandle)
	mux.HandleFunc("/azhar", handler.AzharHandle)
	mux.HandleFunc("/product", handler.ProductHandle)
	mux.HandleFunc("/post-get", handler.PostGet)
	mux.HandleFunc("/form", handler.Form)
	mux.HandleFunc("/process", handler.Process)

	fileServer := http.FileServer(http.Dir("assets"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Println("Start on port 8080")

	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}
