package main

import "fmt"

type playerchar struct {
	name                     string
	STR                      uint8
	DEX                      uint8
	CON                      uint8
	INT                      uint8
	WIS                      uint8
	CHA                      uint8
	nextExpToLevel           uint64
	currentExp               uint64
	inventoryMaxKilograms    float64
	inventorCurrentKilograms float64
}

func (pc *playerchar) displayPlayerCharacterStats() {
	fmt.Println("Your character's name is:", pc.name)
	fmt.Println("========================================")
	fmt.Println("STR:", pc.STR)
	fmt.Println("DEX:", pc.DEX)
	fmt.Println("CON:", pc.CON)
	fmt.Println("INT:", pc.INT)
	fmt.Println("WIS:", pc.WIS)
	fmt.Println("CHA:", pc.CHA)
	fmt.Println("========================================")
}

func main() {
	lysander := playerchar{"Lysander", 8, 9, 5, 12, 8, 19, 1490, 1129, 38.2, 31.0}
	lysander.displayPlayerCharacterStats()
}
