package main

import (
	"fmt"
	"net/http"
)

type grilledcheese struct {
	name     string
	bread    string
	cheese   string
	fillings []string
	price    float32
}

func (r *grilledcheese) readGrilledcheese() {
	fmt.Println("Here's what in a", r.name)
	fmt.Println(r.bread)
	fmt.Println(r.cheese)
	fmt.Println(r.fillings)
	fmt.Println(r.price)
}

func (r *grilledcheese) changeCheese() {
	fmt.Scan(&r.cheese)
	fmt.Println("you have chosen new cheese:", r.cheese)
}

var crabbymelt = grilledcheese{"crabby melt", "sourdough", "mozzerella", []string{"crab meat", "old bay", "scallions"}, 8.99}

func displaygrilledcheese(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "%+v\n", crabbymelt) //takes crabbymelt, formats it to verb %v, then prints to the http.ResponseWriter io stream
}

func main() {
	fmt.Println("We will be listing some of the most popular grilled cheeses")
	//var Crabbymelt = grilledcheese{"crabby melt", "sourdough", "mozzerella", []string{"crab meat", "old bay", "scallions"}, 8.99}
	crabbymelt.readGrilledcheese()
	//crabbymelt.changeCheese()

	http.Handle("/", http.FileServer(http.Dir("./")))
	http.HandleFunc("/grilledcheese/", displaygrilledcheese)
	panic(http.ListenAndServe(":8000", nil))
}
