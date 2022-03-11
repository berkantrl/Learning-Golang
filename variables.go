package main

import (
	"fmt"
	"strconv"
)

var i int = 50

func main() {
	println(i)
	//var i int
	//i = 42
	//i = 27
	//j := 23
	// fmt.Println(i)
	// fmt.Printf("%v, %T", j, j)
	// println(j)

	var (
		actorName string = "Johnny Deep"
		companion string = "Leonardo Da Vinci"
		age       int    = 52
		movieDay  int    = 21
	)
	println(actorName, companion, age, movieDay)

	// Type Conversion
	var i float32 = 42.5
	fmt.Printf("%v, %T\n", i, i)

	var j int
	j = int(i)
	fmt.Printf("%v, %T\n", j, j)

	// integer to string
	// use strconv package

	var number int = 42
	fmt.Printf("%v, %T\n", number, number)

	var name string
	name = strconv.Itoa(number)
	fmt.Printf("%v, %T\n", name, name)

}
