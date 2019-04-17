package main

import(
	"os"
	"flag"
)

func main(){
	mydir := flag.String("directory","test","input a directory you'd like to use")
	flag.Parse()
	
	
	os.Mkdir(*mydir, 0777)
}