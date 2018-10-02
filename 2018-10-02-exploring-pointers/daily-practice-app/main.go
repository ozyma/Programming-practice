package main

import "fmt"

type grilledcheese struct {
	bread    string
	cheese   string
	fillings []string
	price    float32
}

func main() {
	fmt.Println("We will be listing some of the most popular grilled cheeses")
	crabbymelt := grilledcheese{"sourdough", "mozzerella", []string{"crab meat", "old bay", "scallions"}, 8.99}
	fmt.Println(crabbymelt)
	fmt.Println(crabbymelt.bread)
	fmt.Println(crabbymelt.cheese)
	fmt.Println(crabbymelt.fillings)
	fmt.Println(crabbymelt.price)

}
