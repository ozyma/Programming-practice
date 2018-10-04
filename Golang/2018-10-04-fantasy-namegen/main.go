package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// creating a map of prefixes where int is the key, and a string in the value
	prefix := make(map[int]string)
	prefix[0] = "Oa"
	prefix[1] = "Vy"
	prefix[2] = "Ly"
	prefix[3] = "Ge"
	prefix[4] = "Ji"
	prefix[5] = "Oi"
	prefix[6] = "Qe"
	prefix[7] = "Xe"
	prefix[8] = "He"
	prefix[9] = "Ku"

	// creating a map of middle characters where int is the key, and a string in the value
	middle := make(map[int]string)
	middle[0] = "bart"
	middle[1] = "vak"
	middle[2] = "lus"
	middle[3] = "paw"
	middle[4] = "xuth"
	middle[5] = "req"
	middle[6] = "huk"
	middle[7] = "lut"
	middle[8] = "weq"
	middle[9] = "cuv"

	// creating a map of suffixes where int is the key, and a string in the value
	suffix := make(map[int]string)
	suffix[0] = "ert"
	suffix[1] = "us"
	suffix[2] = "axy"
	suffix[3] = "uva"
	suffix[4] = "ytyr"
	suffix[5] = "uyto"
	suffix[6] = "aak"
	suffix[7] = "uug"
	suffix[8] = "yyk"
	suffix[9] = "eem"

	// creating a map of titles where int is the key, and a string in the value
	title := make(map[int]string)
	title[0] = "Ranger"
	title[1] = "Psion"
	title[2] = "Android"
	title[3] = "Repairman"
	title[4] = "Mechanic"
	title[5] = "Warden"
	title[6] = "Bandit"
	title[7] = "Hunter"
	title[8] = "Assassin"
	title[9] = "Hypo Junkie"

	//Randomly generate names 10 times
	for i := 0; i < 100; i++ {

		//makes a new random number source based upon the current nanosecond value since Jan. 1 1970 UTC
		randnumbersource := rand.NewSource(time.Now().UnixNano())
		//the random func needs what's called a "source," or a number to base its randomness off. We use the nanosecond value about in order to base our randomness
		randnumber := rand.New(randnumbersource)
		randprefix := prefix[randnumber.Intn(10)]
		randmiddle := middle[randnumber.Intn(10)]
		randsuffix := suffix[randnumber.Intn(10)]
		randtitle := title[randnumber.Intn(10)]

		//using %v verb for Printf in order to format my string varliables correctly
		fmt.Printf("%v%v%v, the %v", randprefix, randmiddle, randsuffix, randtitle)
		fmt.Println()
		//because the program is too fast for the UnixNano() function, we wait 2 milliseconds between each loop call
		time.Sleep(2 * time.Millisecond)
	}
}
