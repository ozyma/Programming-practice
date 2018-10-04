// Writing files in Go follows similar patterns to the
// ones we saw earlier for reading.

package main

import (
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// To start, here's how to dump a string (or just
	// bytes) into a file.
	d1 := []byte("hello\ngo\n")
	//err := ioutil.WriteFile("/tmp/dat1", d1, 0644)
	//check(err)

	// For more granular writes, open a file for writing.
	f, err := os.Create("./tmp/dat2")
	err = ioutil.WriteFile("./tmp/dat2", d1, 0644)
	check(err)

	f.Close()

}
