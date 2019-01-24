package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Page struct {
	Name string
}

func main() {
	
	template := template.Must(template.ParseFiles("templates/index.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		//присвоение имени
		p := Page{Name: "Gopher"}


		if name := r.FormValue("name"); name != "" {
			p.Name = name
		}


		//server
		if err := template.ExecuteTemplate(w, "index.html", p); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	})

	fmt.Println(http.ListenAndServe(":8080", nil))
}
