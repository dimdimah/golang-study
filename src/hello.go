package main

import "fmt"

func main() {
	const nama = "Dimah"
	var name = "Luthfi"
	var n uint8 = name[3]
	var nString = string(n)

	fmt.Println(name)
	fmt.Println(n)
	fmt.Println(nString)

	fmt.Println("Nama saya adalah", nama)

	var x = 12
	x += 2
	fmt.Println(x)

	x += 4
	fmt.Println(x)

}
