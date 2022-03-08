package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	// var arr [3]int

	// arr[0] = 1
	// arr[1] = 2
	// arr[2] = 3

	// arr2 := [3]int{1, 2, 3}

	// fmt.Println(arr, arr2)

	// slice := []int{1, 3, 5}

	// println(slice[1])

	// slice2 := arr2[1:]
	// println(slice2[1])

	// m := map[string]int{"orhun": 42}
	// fmt.Println(m)
	// println(m["orhun"])

	var food []string

	food = append(food, "meat")

	println(food[0])

	city := map[string]int{}

	city["Elazığ"] = 23

	println(city["Elazığ"])

}
