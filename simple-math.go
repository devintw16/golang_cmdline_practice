package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if (os.Args[2] == "+"){
		fmt.Println(add())
	} else if ( os.Args[2] == "-"){
		fmt.Println(subtract())
	}
}

func add()(z int){
	fmt.Printf("Adding %s and %s\n",os.Args[1],os.Args[3])
	z = strtoint(os.Args[1]) + strtoint(os.Args[3])
	return
}

func subtract()(z int){
	fmt.Printf("Subtracting %s from %s\n",os.Args[3],os.Args[1])
	z = strtoint(os.Args[1]) - strtoint(os.Args[3])
	return
}

func strtoint(arg string)(z int){
	var err error
	if z, err = strconv.Atoi(arg); err != nil {
		panic(err)
	}
	return
}