package main

import(
	"os"
)

func main(){
	os.Mkdir("asset", 0777)
	os.Mkdir("src", 0777)
	os.Mkdir("third-party", 0777)
}