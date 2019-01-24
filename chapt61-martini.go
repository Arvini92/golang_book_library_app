package main

import (
	"github.com/go-martini/martini"
	"net/http"
)

type Person struct {
	Name string
}

func setName1(c martini.Context, req *http.Request) {
	c.Map(Person{Name: "Gopher"})
}

func setName2(c martini.Context, req *http.Request) {
	c.Map(Person{Name: "Golfer"})
}

func index(p Person) string {
	return "Hello, " + p.Name + "\n"
}

func getBooks(p Person) string {
	return "Getting all books, " + p.Name + "\n"
}

func getBook(params martini.Params) string {
	return "Getting book with ID" + params["id"] + "\n"
}

func deleteBook(params martini.Params) string {
	return "Deleting book with ID" + params["id"] + "\n"
}

func main() {
	m := martini.Classic()

	//index
	m.Get("/", index)
	m.Use(setName1)

	//books
	m.Group("/books", func(r martini.Router) {
		r.Get("/", getBooks)
		r.Get("/:id", getBook)
		r.Delete("/:id", deleteBook)
	}, setName2)

	// serve on port 3000
	m.Run()

}
