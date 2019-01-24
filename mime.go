package main

import (
    "fmt"
    "net/http"
    "sort"
    _"log"
)

func handler(w http.ResponseWriter, r *http.Request) {

    var keys []string
    for k := range r.Header {
        keys = append(keys, k)
    }
    sort.Strings(keys)
    
    fmt.Fprintln(w, "<b>Request Headers:</b></br>", r.URL.Path[1:])
    for _, k := range keys {
        fmt.Fprintln(w, k, ":", r.Header[k], "</br>", r.URL.Path[1:])
        fmt.Println(r.Header[k])
    }
    //log.Println(w, r.Header["Content-Type"])
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}

logEntry := "Content-Encoding: gzip\r\nLast-Modified: Tue, 20 Aug 2013 15:45:41 GMT\r\nServer: nginx/0.8.54\r\nAge: 18884\r\nVary: Accept-Encoding\r\nContent-Type: text/html\r\nCache-Control: max-age=864000, public\r\nX-UA-Compatible: IE=Edge,chrome=1\r\nTiming-Allow-Origin: *\r\nContent-Length: 14888\r\nExpires: Mon, 31 Mar 2014 06:45:15 GMT\r\n"

// don't forget to make certain the headers end with a second "\r\n"
reader := bufio.NewReader(strings.NewReader(logEntry + "\r\n"))
tp := textproto.NewReader(reader)

mimeHeader, err := tp.ReadMIMEHeader()
if err != nil {
    log.Fatal(err)
}

// http.Header and textproto.MIMEHeader are both just a map[string][]string
httpHeader := http.Header(mimeHeader)
log.Println(httpHeader)

logEntry := "Content-Encoding: gzip\r\nLast-Modified: Tue, 20 Aug 2013 15:45:41 GMT\r\nServer: nginx/0.8.54\r\nAge: 18884\r\nVary: Accept-Encoding\r\nContent-Type: text/html\r\nCache-Control: max-age=864000, public\r\nX-UA-Compatible: IE=Edge,chrome=1\r\nTiming-Allow-Origin: *\r\nContent-Length: 14888\r\nExpires: Mon, 31 Mar 2014 06:45:15 GMT\r\n"

// we need to make sure to add a fake HTTP header here to make a valid request.
reader := bufio.NewReader(strings.NewReader("GET / HTTP/1.1\r\n" + logEntry + "\r\n"))

logReq, err := http.ReadRequest(reader)
if err != nil {
    log.Fatal(err)
}

log.Println(logReq.Header)