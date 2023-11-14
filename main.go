package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Jugador struct {
	Apodo  string
	Equipo string
	Region string
	Elo    string
}

func main() {
	mux := mux.NewRouter()

	mux.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ahora := time.Now()

		tmpl := template.Must(template.ParseFiles("templates/index.html"))
		tmpl.Execute(w, "")

		fmt.Println("Buenas tardes | Tomo:", time.Since(ahora))
	}).Methods("GET")

	http.ListenAndServe(":8080", mux)
}
