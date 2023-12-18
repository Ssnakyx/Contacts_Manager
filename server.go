package main

import (
	"log"
	"net/http"
	"text/template"
)

type Informations struct {
	FirstName   string
	LastName    string
	Id          int
	PhoneNumber int
	Status      string
	Email       string
}

func Home(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("HTML/index.html", "HTML/info.html")
	if err != nil {
		log.Fatal(err)
	}
	template.Execute(w, nil)
}

func FormulaireHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Erreur", http.StatusInternalServerError)
			return
		}
	}
	http.Redirect(w, r, "http://localhost:8084", http.StatusFound)
}

func Infos(w http.ResponseWriter, r *http.Request, infos *Informations) {

	template, err := template.ParseFiles("HTML/index.html", "HTML/info.html")
	if err != nil {
		log.Fatal(err)
	}
	template.Execute(w, infos)
}

func main() {
	//infos := &Informations("H")
	http.HandleFunc("/", Home)
	//http.HandleFunc("/infos", func(w http.ResponseWriter, r *http.Request) {
	//Infos(w, r, infos)
	//})
	fs := http.FileServer(http.Dir("CSS/"))
	http.Handle("/CSS/", http.StripPrefix("/CSS", fs))
	log.Println("Serveur allumé")
	http.ListenAndServe(":8084", nil)
}

// Je l'ai fait avec 8084 car avec 8080 j'avais un conflit avec mon Hangman Web
