package main

import (
	"net/http"
	"log"
	"html/template"
)

func main() {
	fs := http.FileServer(http.Dir("view/static/"))
	http.Handle("/static/", http.StripPrefix("/static/",fs))
	http.HandleFunc("/", seveIndex)
	log.Println("Listening...")
	http.ListenAndServe(":8080", nil)
}
func seveIndex(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("view/index.html"))
	tmpl.Execute(w, nil)
}