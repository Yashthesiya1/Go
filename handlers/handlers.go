package handlers

import (
	"html/template"
	"net/http"
)

type dataStore struct {
	Name    string
	Amount  int
	Profile string
	Quote   string
}

var Templates *template.Template

func ReviewHandler(response http.ResponseWriter, request *http.Request) {
	data := dataStore{
		Name:    "Yash Thesiya",
		Amount:  85,
		Profile: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSeE8Sti7nupldscJu8jVr8ecoimhS5tkG-3Q&s",
		Quote:   "Donec ultricies arcu quis nisl pellentesque, quis lobortis augue. Donec ultricies arcu quis nisl pellentesque, quis lobortis augue. Donec ultricies arcu quis nisl pellentesque, quis lobortis augue.",
	}
	Templates.ExecuteTemplate(response, "review.html", data)
}

func AcceptedHandle(response http.ResponseWriter, request *http.Request) {
	Templates.ExecuteTemplate(response, "accepted.html", nil)
}
func DeniedHandle(response http.ResponseWriter, request *http.Request) {
	Templates.ExecuteTemplate(response, "denied.html", nil)
}
