package main

import (
	"fmt"
	"net/http"
	"image"
	"image/jpeg"
    "encoding/json"
	"os"
	
	"io"
	"database/sql"
    _ "github.com/mattn/go-sqlite3"
	"math/rand"
	
)

type Parent struct {
	ID string
	Title string
	Image string
}

type ParentIns struct {
	Title string `json:"title"`
	Image string `json:"image"`
}
func indexAdd(w http.ResponseWriter, r *http.Request) {

	if origin := r.Header.Get("Origin"); origin != "" {
        w.Header().Set("Access-Control-Allow-Origin", origin)
        fmt.Println(origin)
    }
    w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token")

	var parent ParentIns
	randName := rand.Intn(1000000)
	imageName := "image/" + fmt.Sprintf("%d", randName) + ".jpg"

	 if r.Method == "POST" {

		headerType := r.Header["Content-Type"]
		imageHeader := headerType[0]
		imageType :="image/jpeg"

		if imageHeader == imageType {				
			img, _, _ := image.Decode(r.Body)
			
			
			file, _ := os.Create("./" + imageName)
			jpeg.Encode(file, img, &jpeg.Options{100})
			
		}else{


			dec := json.NewDecoder(r.Body)
			for {
				if err := dec.Decode(&parent); err == io.EOF {
					break
				} 
			}
			
			db, _ := sql.Open("sqlite3", "./project.db")
			stmt, _ := db.Prepare("INSERT INTO parent(title, image) values(?, ?)")
			stmt.Exec(parent.Title, imageName)
			
			db.Close()

			fmt.Println(parent.Title)
			fmt.Println(imageName)

			//http.Redirect(w, r, "/", 302)
		}
		

		}	
	}


	func main() {

    //mux := http.NewServeMux()
	//mux.HandleFunc("/", indexAdd)
    http.HandleFunc("/add", indexAdd)
    
    //n := negroni.Classic()
	//n.Use(negroni.HandlerFunc(verifyDateBase))
	//n.UseHandler(mux)
	//n.Run(":7070")    

	fmt.Println(http.ListenAndServe(":7070", nil))


	
}