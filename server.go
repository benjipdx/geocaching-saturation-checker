//simple server example in go
//from https://github.com/ljgww/web_server_example_in_Go_-golang-/blob/master/ws_u.go
//http://golang.org/pkg/net/http/#FileServer

package main

import "net/http"
import "html/template"
//import "io"
//import "io/ioutil"

var uploadTemplate, _ = template.ParseFiles("map.html")

func handler(w http.ResponseWriter, r *http.Request) {
    uploadTemplate.Execute(w, nil)
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
