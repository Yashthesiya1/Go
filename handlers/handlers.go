package handlers

import (
	"html/template"
	"net/http"
)

var Templates *template.Template

func ReviewHandler(response http.ResponseWriter, request *http.Request) {
	data := map[string]string{
		"Name":    "Yash Thesiya",
		"Amount":  "60",
		"Profile": "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSeE8Sti7nupldscJu8jVr8ecoimhS5tkG-3Q&s",
	}
	Templates.ExecuteTemplate(response, "review.html", data)
}

func AcceptedHandle(response http.ResponseWriter, request *http.Request) {
	Templates.ExecuteTemplate(response, "accepted.html", nil)
}
func DeniedHandle(response http.ResponseWriter, request *http.Request) {
	Templates.ExecuteTemplate(response, "denied.html", nil)
}
