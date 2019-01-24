	file, err := os.Create("./" + fhead.Filename + "." + ext)
		checkError(err)
		defer file.Close()

		ln, err := io.WriteString(file, content)
		checkError(err)


        
	"io"
    "io/ioutil"
    "os"	

    func checkError(err error) {
    if err != nil {
        panic (err)
    }

    w.Header().Set("Content-type", "image/jpeg")
		w.Header().Set("Content-Disposition", "filename=\"" + parent.Image + "\"")
		jpeg.Encode(w, img, &jpeg.Options{100})

         {{if .Image}}
          {{end}}

          <div><img src="4555.jpg"/></div>



          w.Header().Set("Content-type", "text/html")
	t, _ := template.ParseFiles("index.html")
	t.Execute(w, &page{Title: "Convert Image"})

	

	if r.FormValue("add") != "" {


        	t.Execute(w, parent)


    func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/plain")
	w.Write([]byte("Hello World!!!"))
}


rows, _ := db.Query("select pk,title,author,classification from books")
		for rows.Next() {
			var b Book
			rows.Scan(&b.PK, &b.Title, &b.Author, &b.Classification)
			p.Books = append(p.Books, b)
		}



        rows, _ := db.Query("SELECT id,title,image FROM parent")
		for rows.Next() {
			var parent Parent
			rows.Scan(id, title, image)
            results = append(results, parent)
		}

        rows.Close()


func handler(w http.ResponseWriter, r *http.Request) {
     if r.Method == "POST" {
           r.ParseForm()
           // они все тут
           params := r.Form
     }
}

type Request struct {
    ...
    // Тут данные из query, url, и post - все вместе.
    // Пусто пока не вызовешь ParseForm()
    Form url.Values

    // Содержит данный форм из POST, PATCH или PUT.
    // Пусто пока не вызовешь ParseForm()
    PostForm url.Values

    // Содержит multipart form, включая загрузку файлов.
    // Пусто пока не вызовешь MultipartForm()
    MultipartForm *multipart.Form
    ...
}

