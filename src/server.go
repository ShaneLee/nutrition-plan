package main

import (
	"net/http"
	"text/template"
)

var tmpl = template.Must(template.ParseGlob("public/*"))

func Index(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "Index", nil)
}
