package main

import (
    "fmt"
//    "net"
    "net/http"
//    "net/http/fcgi"
    "html/template"
    "log"
)

func apiHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "{\"test\":true}")
}
func indexHandler(w http.ResponseWriter, r *http.Request) {
    t := template.Must(template.ParseFiles("/go/src/templates/index.html.tpl"))
    str := "Sample Message"
    if err := t.ExecuteTemplate(w, "index.html.tpl", str); err != nil {
        log.Fatal(err)
    }
}

func main() {

    http.HandleFunc("/api", apiHandler)
    http.HandleFunc("/index", indexHandler)
    http.ListenAndServe(":9000", nil)


    // l, err := net.Listen("tcp", "127.0.0.1:9000")
    // if err != nil {
    //     return
    // }
    // http.HandleFunc("/api", apiHandler)
    // http.HandleFunc("/index", indexHandler)
    // fcgi.Serve(l, nil)
}
