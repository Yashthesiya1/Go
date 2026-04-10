package main

import (
	"html/template"
	"log"
	"myFristapp/handlers"
	"net/http"
)

func main() {

	handlers.Templates = template.Must(template.ParseGlob("templates/*.html"))

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", handlers.ReviewHandler)
	http.HandleFunc("/accepted", handlers.AcceptedHandle)
	http.HandleFunc("/denied", handlers.DeniedHandle)
	log.Println("Server started on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
