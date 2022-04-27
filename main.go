package main

import (
	"html/template"
	"log"
	"net/http"
)

type Firstpage struct {
	Valeur string
}

func personnages(w http.ResponseWriter, r *http.Request) {

	template, _ := template.ParseFiles("personnages.html")

	template.Execute(w, r)

}

func main() {

	css := http.FileServer(http.Dir("./css"))
	http.Handle("/css/", http.StripPrefix("/css/", css))
	http.HandleFunc("/", homepage)
	http.HandleFunc("/perso", personnages)
	log.Fatal(http.ListenAndServe(":88", nil))
}
func homepage(w http.ResponseWriter, r *http.Request) {

	template3, _ := template.ParseFiles("ydays.html")
	template3.Execute(w, r)

}
