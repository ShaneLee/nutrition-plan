package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", Index)
	http.ListenAndServe(":7070", nil)
}
