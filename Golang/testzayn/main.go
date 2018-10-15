package main

import (
	"fmt"
	"time"
)

func helpmeprint(s string) {
	for i := 0; i <= 100; i++ {
		fmt.Println(s)
		time.Sleep(2 * time.Millisecond)
	}

}

func main() {
	go helpmeprint("Tanner")
	go helpmeprint("Zayn")
	go helpmeprint("Shayon")
	fmt.Scanln()

}
