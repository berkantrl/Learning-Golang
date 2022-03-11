package main

import "fmt"

func main() {

	fmt.Println("Hello World")

	//Boolean

	// var n bool = false
	// fmt.Printf("%v, %T", n, n)

	test := 1 == 1
	test2 := 1 == 2
	fmt.Printf("%v, %T ", test, test)
	fmt.Printf("%v, %T ", test2, test2)

	// uint8 -- uint16 -- uint32
	n := 5
	fmt.Printf("\n%v, %T\n", n, n)
	var a uint16 = 42
	fmt.Printf("%v, %T\n", a, a)

	var i int = 4
	var j int8 = 2
	fmt.Println(i + int(j))

	b := 8
	fmt.Println(b << 5)
	fmt.Println(b >> 5)

	// Float 32 -- Float64
	// s := 3.14
	// f := 13.7e72
	// d := 2.1e14
	// fmt.Printf("%v, %T", s, s)
	// fmt.Printf("\n%v, %T", f, f)
	// fmt.Printf("\n%v, %T", d, d)

	var s float32 = 3.14
	var f float64 = 13.7e72
	var d float32 = 2.1e14
	fmt.Printf("%v, %T", s, s)
	fmt.Printf("\n%v, %T", f, f)
	fmt.Printf("\n%v, %T", d, d)

	// complex64 -- complex128
	var y complex64 = 2i
	fmt.Printf("\n%v, %T", y, y)

	var g complex128 = complex(5, 12)
	fmt.Printf("\n%v, %T", g, g)

	// string
	p := "String example"
	fmt.Println(string(p[2]))
	bt := []byte(p)
	fmt.Println(bt)
	var r rune = 'a'
	fmt.Printf("%v, %T", r, r)

}
