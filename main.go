package main

import (
	"fmt"
)

// import (
// 	"fmt"
// 	"html/template"
// 	"net/http"
// )

// type Page struct {
// 	Valeur string
// }

// var templates = template.Must(template.ParseFiles("index.html"))

// func main() {
// 	http.HandleFunc("/", homeHandler)
// 	fmt.Println("Server Running")
// 	http.ListenAndServe(":9999", nil)
// }

// func homeHandler(w http.ResponseWriter, r *http.Request) {
// 	p := Page{Valeur: "mon premier essai"}
// 	err := templates.ExecuteTemplate(w, "index.html", p)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 	}
// }

func main() {
	fmt.Println("Hello World")
}
