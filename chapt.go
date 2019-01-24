package main

import (
	"fmt"
	"net/http"
	"image"
	"image/jpeg"
    
	"os"
	

	"database/sql"
    _ "github.com/mattn/go-sqlite3"

    _"github.com/codegangsta/negroni"
	
)


type ParentIns struct {
	Title string
	Image string
}

type Parent struct {
	ID string
	Title string
	Image string
}




func indexAdd(w http.ResponseWriter, r *http.Request) {
        if r.Method == "POST" {
		imgfile, fhead, _ := r.FormFile("image")
        
		img, _, _ := image.Decode(imgfile)

		file, _ := os.Create("./image/" + fhead.Filename)
		jpeg.Encode(file, img, &jpeg.Options{100})

		

		var parent ParentIns
		parent.Title = r.FormValue("title") 
		parent.Image = ("image/" + fhead.Filename)

		db, _ := sql.Open("sqlite3", "./project.db")
		stmt, _ := db.Prepare("INSERT INTO parent(title, image) values(?,?)")
		stmt.Exec(parent.Title, parent.Image)
        //db.Close()

		fmt.Println(parent.Title)
		fmt.Println(parent.Image)

    	//http.Redirect(w, r, "/", 302)
        }else{
            
        }
	}

  



func main() {

    //mux := http.NewServeMux()
	//mux.HandleFunc("/", indexAdd)
    http.HandleFunc("/", indexAdd)
    
    //n := negroni.Classic()
	//n.Use(negroni.HandlerFunc(verifyDateBase))
	//n.UseHandler(mux)
	//n.Run(":7070")    

	//fmt.Println(http.ListenAndServe(":7070", nil))


	
}














		
