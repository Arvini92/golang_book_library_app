package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/codegangsta/negroni"
	_ "github.com/mattn/go-sqlite3"
	"image"
	"image/jpeg"
	"io"
	"math/rand"
	"net/http"
	"os"
	"strconv"
)

type ParentIns struct {
	Title string `json:"title"`
	Image string `json:"image"`
}

type MyParent struct {
	ID    string
	Title string
	Image string
}


type Child struct {
		IdChild string
		TitleChild string
		ImageChild string
		IdParent string
	}


type Parent struct {
	ID    string
	Title string
	Image string
	Children []Child
}

type All struct{
	ID    string
	Title string
	Image string
	IdChild string
	TitleChild string
	ImageChild string
	IdParent string
}
func indexHandler(w http.ResponseWriter, r *http.Request) {

	if origin := r.Header.Get("Origin"); origin != "" {
		w.Header().Set("Access-Control-Allow-Origin", origin)
		fmt.Println(origin)
	}
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token")

	db, _ := sql.Open("sqlite3", "project.db")

	myparents := []MyParent{}

	rows, _ := db.Query("SELECT * FROM parent")
	childs, _ := db.Query("SELECT id_child, title_child, image_child, id_parent FROM parent INNER JOIN child ON parent.id = child.id_parent")
	alls, _ := db.Query("SELECT * FROM parent INNER JOIN child ON parent.id = child.id_parent")
	rowsi, _ := db.Query("SELECT * FROM parent")

	for rows.Next() {
		var myparent MyParent
		rows.Scan(&myparent.ID, &myparent.Title, &myparent.Image)
		myparents = append(myparents, myparent)
	}
	mychildren := []Child{}

	for childs.Next() {
		var child Child
		childs.Scan(&child.IdChild, &child.TitleChild, &child.ImageChild, &child.IdParent)

		mychildren = append(mychildren, child)
	}

	parents := []Parent{}
	for rowsi.Next() {
		var parent Parent
		rowsi.Scan(&parent.ID, &parent.Title, &parent.Image)
		fmt.Println(parent.ID)
		clid, _ := db.Query("SELECT *FROM child WHERE id_parent = ?", &parent.ID)	
		fmt.Println(clid)
		
		for clid.Next(){
		var child Child
		clid.Scan(&child.IdChild, &child.TitleChild, &child.ImageChild, &child.IdParent)
		parent.Children = append(parent.Children, child)
		}
		parents = append(parents, parent)
	}



	myAll := []All{}
	for alls.Next() {
		var all All
		alls.Scan(&all.ID, &all.Title, &all.Image, &all.IdChild, &all.TitleChild, &all.ImageChild, &all.IdParent)

		myAll = append(myAll, all)
	}

	//rows.Close()

	//Write slected items
	json.NewEncoder(w).Encode(parents)

	if r.Method == "DELETE" {

		title := r.URL.Path[len("/"):]

		if title != "" {
			id, _ := strconv.ParseInt(title, 10, 64)
			fmt.Println(title)
			fmt.Println("DELETE")
			db, _ := sql.Open("sqlite3", "project.db")
			db.Exec("delete from parent where id = ?", id)
			//db.Close()

		}

	}

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
		imageType := "image/jpeg"

		if imageHeader == imageType {
			img, _, _ := image.Decode(r.Body)

			file, _ := os.Create("./" + imageName)
			jpeg.Encode(file, img, &jpeg.Options{100})

		} else {

			dec := json.NewDecoder(r.Body)
			for {
				if err := dec.Decode(&parent); err == io.EOF {
					break
				}
			}

			db, _ := sql.Open("sqlite3", "./project.db")
			stmt, _ := db.Prepare("INSERT INTO parent(title, image) values(?, ?)")
			stmt.Exec(parent.Title, imageName)

			//db.Close()

			fmt.Println(parent.Title)
			fmt.Println(imageName)

			//http.Redirect(w, r, "/", 302)
		}

	}
}

func main() {

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/add", indexAdd)
	fmt.Println(http.ListenAndServe(":7070", nil))

}
