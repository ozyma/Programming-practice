package main

import (
	"fmt"
	"net/http"
)

func indexhandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Go is neat!") //Fprintf is going to use "w" and format "Go is neat!" to it, therefore printing it onto our http response.
}

func main() {
	http.HandleFunc("/", indexhandler)
	http.ListenAndServe(":8080", nil) //nil will default us to the Go const DefaultMuxServer with is used by http package by default.

}
