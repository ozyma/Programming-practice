package main

import (
	"fmt"
)

type car struct {
	name  string
	speed int32
	cost  int32
}

func main() {
	var (
		x, y, z int
	)
	corvette := car{"corvette", 32, 26000}
	fmt.Println(corvette)
	fmt.Println(corvette.name)
	fmt.Println(corvette.cost)
	fmt.Println(corvette.speed)
	fmt.Println("hello test!")
	fmt.Println(x, y, z)
	fmt.Scanln()

}
