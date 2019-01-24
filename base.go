package main
import(
"fmt"
"database/sql"
_ "github.com/mattn/go-sqlite3"
)

func main() {

	db, _ := sql.Open("sqlite3", "project.db")
    chlids, _ := db.Query("SELECT id, title FROM parent INNER JOIN child ON parent.id = child.id_parent")
    var a, b, c, d, f, g, h string
    for chlids.Next(){
    
    chlids.Scan(a, b, c, d, f, g, h)  
    fmt.Println(a)
    fmt.Println(b)
    fmt.Println(c)
    fmt.Println(d)
    fmt.Println(f)
    fmt.Println(g)
    fmt.Println(h)
    }
}