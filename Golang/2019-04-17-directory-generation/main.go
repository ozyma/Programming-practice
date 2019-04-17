package main

import(
	"os"
)

func main(){
	os.Mkdir("test/", 0777)
}