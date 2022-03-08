package main

import "fmt"

// harf büyük olursa public oluyor küçük harfle başlıyorsa private func oluyor
func main() {
	 fmt.Println("Hello Security !")

	 var i int
	 i = 42
	 println(i)

	 var f float32 = 3.14
	 println(f)

	 name := "Berkan"
	 print(name)

	 println(i, f, name)

	 var lastname *string = new(string)
	 *lastname = "Türel"

	 fmt.Println(lastname, *lastname)

	 var pname *string = &name

	 *pname = "Yousuf"

	 println(name)

	 const age int = 20

	 println(age)

	 println(first, second, thirty)
	 println(monday, tday, wday)

}

const (
 	first = iota
	second
 	thirty
)

const (
	monday = iota + 1
 	tday   = 2 << iota
	wday
)
