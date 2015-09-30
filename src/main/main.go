// main
package main

import (
//	"fmt"
	"net/http"
	"log"
//	"html"
)

func main() {
	startServer()
}

func ezzahHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("form ", r.Form, " body", r.Body, " url ", r.URL," postVal(\"user\") ", r.PostFormValue("user"))
	w.Write([]byte("<h1>Ezzah</h1>"))
}

func startServer() {
	http.HandleFunc("/bar", ezzahHandler)
	log.Fatal(http.ListenAndServe(":8090", nil))
}

