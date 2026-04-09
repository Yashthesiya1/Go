package handlers

import (
	"html/template"
	"net/http"
)

var Templates *template.Template

func ReviewHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"Name":   "Arif Momin",
		"Amount": "60",
	}
	Templates.ExecuteTemplate(w, "review.html", data)
}
