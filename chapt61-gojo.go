package main

import (
	"fmt"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
	"net/http"
)

func setName(c *web.C, h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		c.Env["name"] = "Gopher"
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

func index(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!\n", c.Env["name"])
}

func getBooks(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Getting all the boooks\n")
}

func getBook(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Getting book with ID %s\n", c.URLParams["id"])
}

func deleteBook(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Deleting book with ID %s\n", c.URLParams["id"])
}

func main() {
	//index
	goji.Get("/", index)

	//books
	goji.Get("/books", getBooks)

	// invividual books
	goji.Get("/books/:id", getBook)
	goji.Delete("/books/:id", deleteBook)

	// middleware
	goji.Use(setName)

	//serve on port 8080
	goji.Serve()
}
