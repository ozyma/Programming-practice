package main

import "fmt"

func main() {
	var primus int32 = 45      //declare an int32 value of 45 for variable primus
	primusmemory := &primus    //the value of primusmemory is the memory address of primus' value. 0xc00004c058 for this instance
	fmt.Println(primus)        //print the value of variable
	fmt.Println(&primus)       //print the memory address of primus
	fmt.Println(*primusmemory) //point to whatever's being stored at the memory address 0xc00004c058, print the value
}
