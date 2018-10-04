# October 4th, 2018 - Fantasy struct praactice

I am going to throw out struct ideas/methods for my PDM rogue-like for fun and practice

I was unable to use this function earlier:
```go
func main() {
	lysander := playerchar{"Lysander", 8, 9, 5, 12, 8, 19, 1490, 1129, 38.2, 31.0}
	displayPlayerCharacterStats(lysander)
}
```

And I realized it was because my function:
```go
func (pc *playerchar) displayPlayerCharacterStats() {
    fmt.Println("Your character's name is:", pc.name)
}
```

was a method for playerchar types, and therefore has to be called as such:
```go
lysander.displayPlayerCharacterStats()
``` 
