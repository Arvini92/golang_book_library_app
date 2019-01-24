package main

import (
	"fmt"
	"html/template"
	"net/http"

	"database/sql"
	_ "github.com/mattn/go-sqlite3"

	"encoding/json"

	
)

type Parent struct {
	ID string
	Title string
	Image string
}




func indexHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates2/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	db, _ := sql.Open("sqlite3", "project.db")


    
    t.ExecuteTemplate(w, "index.html", nil)
 
    
 
    results := []Parent{
			Parent{"1", "Anna", "DSC_0038.jpg"},
			Parent{"2", "Tanya", "DSC_0039.jpg"},
			Parent{"3", "Galya", "DSC_0040.jpg"},
			Parent{"4", "Sveta", "DSC_0041.jpg"},
			Parent{"5", "Rita", "DSC_0042.jpg"},
		}

        
    rows, _ := db.Query("SELECT * FROM parent")
    

    
    for rows.Next() {
        var parent Parent
        rows.Scan(&parent.ID, &parent.Title, &parent.Image)
        
        results = append(results, parent)
    }
       
    rows.Close() 
	
       
    
    json.NewEncoder(w).Encode(results)
	
}



func main() {



	http.HandleFunc("/", indexHandler)


	fmt.Println(http.ListenAndServe(":7070", nil))
}




		