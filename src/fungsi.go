package main

import (
	"fmt"
	"strings"
)

func printMessage(message string, args []string) {
	var nameString = strings.Join(args, " ")
	fmt.Println(message, nameString)
}

func main() {
	var names = []string{"John", "Wick"}
	printMessage("halo", names)
}
