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
	"github.com/nfnt/resize"
	"github.com/rs/cors"
)

type ParentIns struct {
	Title string `json:"title"`
	Image string `json:"image"`
}
type ChildIns struct {
	Title string `json:"title"`
	Image string `json:"image"`
	ParentId string `json:"parentId"`
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


var db *sql.DB

func init() {
   
    db, _ = sql.Open("sqlite3", "project.db")
    
}



func indexHandler(w http.ResponseWriter, r *http.Request) {

	parents := getData()

	json.NewEncoder(w).Encode(parents)

}

func getData() ([]*Parent) {
    
	rows, _ := db.Query("SELECT id, title, image FROM project WHERE parent_id = 0")
	
	parents := make([]*Parent, 0)

	for rows.Next() {
		parent := new(Parent)
		rows.Scan(&parent.ID, &parent.Title, &parent.Image)
		childs, _ := db.Query("SELECT id, title, image, parent_id FROM project WHERE parent_id = ?", &parent.ID)

		for childs.Next(){
			var child Child
			childs.Scan(&child.IdChild, &child.TitleChild, &child.ImageChild, &child.IdParent)
			parent.Children = append(parent.Children, child)
		}
		childs.Close()
		parents = append(parents, parent)
	}
	rows.Close()
    

    return parents
}


func deleteParent(w http.ResponseWriter, r *http.Request) {
		if r.Method == "DELETE" {

		title := r.URL.Path[len("/"):]

		if  title != "child/" {
			title := r.URL.Path[len("/parent/"):]
			id, _ := strconv.ParseInt(title, 10, 64)
			fmt.Println(title)

			deleteData(id)
			
		}
	}
}

func deleteChild(w http.ResponseWriter, r *http.Request) {

			if r.Method == "DELETE" {

			title := r.URL.Path[len("/"):]

			if  title != "parent/" {
			title := r.URL.Path[len("/child/"):]
			id, _ := strconv.ParseInt(title, 10, 64)
			
			deleteData(id)

		}

	}
}


func deleteData(id int64) {

			rows, _ := db.Query("SELECT max(lft), max(rht) FROM project Where id = ?", id)

			var lft, rht string

			for rows.Next() {
				rows.Scan(&lft, &rht)
			}

			db.Exec("delete from project where lft >= ? and rht <= ?", lft, rht)
			db.Exec("update project set rht = rht - (? - ? + 1) where rht > ? and lft < ?", rht, lft, rht, lft)
			db.Exec("update project set lft = lft - (? - ? + 1), rht=rht-(? - ? + 1) where lft > ?", rht, lft, rht, lft, rht)
			db.Close()
			fmt.Println(lft)
			fmt.Println(rht)
			fmt.Println("DELETE")

}




func indexAddParent(w http.ResponseWriter, r *http.Request) {



	var parent ParentIns
	

	if r.Method == "POST" {

		randName := rand.Intn(1000000)
		imageName := fmt.Sprintf("%d", randName) + ".jpg"
		
		headerType := r.Header["Content-Type"]
		imageHeader := headerType[0]
		imageType := "image/jpeg"
		if imageHeader == imageType {
	
			img, _, _ := image.Decode(r.Body)

			file, _ := os.Create("./image/" + imageName)
			file80, _ := os.Create("./image80x80/" + imageName)
			newImage := resize.Resize(80, 80, img, resize.Lanczos3)
			jpeg.Encode(file, img, &jpeg.Options{100})
			jpeg.Encode(file80, newImage, &jpeg.Options{100})
			
		} else {

			dec := json.NewDecoder(r.Body)
			for {
				if err := dec.Decode(&parent); err == io.EOF {
					break
				}
			}
			
			addDataPar(parent.Title, imageName)
					
		}
	}
}


func addDataPar(title string, image string) {

			rows, _ := db.Query("SELECT max(lft), max(rht) FROM project")

			var lft, rht string

			for rows.Next() {
				rows.Scan(&lft, &rht)
			}
			
			stmt, _ := db.Prepare("INSERT INTO `project` (`lvl`, `lft`, `rht`, `parent_id`, `title`, `image` ) VALUES (1, ? + 1, ? + 2 , 0, ?, ?)")
			fmt.Println(lft)
			fmt.Println(rht)
			stmt.Exec(rht, rht, title, image)

			db.Close()
			fmt.Println(title)
			fmt.Println(image)

}


func indexAddChild(w http.ResponseWriter, r *http.Request) {


	var child ChildIns
	
	if r.Method == "POST" {

		randName := rand.Intn(1000000)
		imageName := fmt.Sprintf("%d", randName) + ".jpg"
		
		headerType := r.Header["Content-Type"]
		imageHeader := headerType[0]
		imageType := "image/jpeg"
		if imageHeader == imageType {
	
			img, _, _ := image.Decode(r.Body)

			file, _ := os.Create("./image/" + imageName)
			file80, _ := os.Create("./image80x80/" + imageName)
			newImage := resize.Resize(80, 80, img, resize.Lanczos3)
			jpeg.Encode(file, img, &jpeg.Options{100})
			jpeg.Encode(file80, newImage, &jpeg.Options{100})
			
		} else {

			dec := json.NewDecoder(r.Body)
			for {
				if err := dec.Decode(&child); err == io.EOF {
					break
				}
			}

			addDataCh(child.ParentId, child.Title, imageName)
			
		}
	}
}

func addDataCh(parId string, title string, image string) {

			rows, _ := db.Query("SELECT max(lft), max(rht) FROM project Where id = ?", parId)

			var lft, rht string

			for rows.Next() {
				rows.Scan(&lft, &rht)
			}
			db.Exec("update `project` set `lft` = `lft` + 2, `rht` = `rht` + 2 where `lft` > ?", rht)
			db.Exec("update `project` set `rht` = `rht` + 2 where `rht` = ?", rht)
			stmt, _ := db.Prepare("INSERT INTO `project` (`lvl`, `lft`, `rht`, `parent_id`, `title`, `image`) VALUES (2, ?, ? + 1, ?, ?, ?)")
			fmt.Println(lft)
			fmt.Println(rht)
			stmt.Exec(rht, rht, parId, title, image)
		
			fmt.Println(title)
			fmt.Println(image)
			fmt.Println(parId)

}





func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/child", deleteChild)
	mux.HandleFunc("/parent", deleteParent)
	mux.HandleFunc("/add/child", indexAddChild)
	mux.HandleFunc("/add/parent", indexAddParent)

	
	handler := cors.Default().Handler(mux)
	fmt.Println(http.ListenAndServe(":7070", handler))
	db.Close()
}
