package main

import "fmt"

type playerchar struct {
	name string
	STR  uint8
	DEX  uint8
	CON  uint8
	INT  uint8
	WIS  uint8
	CHA  uint8
}

func main() {
	var lysander = playerchar{"lysander", 8, 9, 5, 12, 8, 19}
	fmt.Println(lysander)

}
