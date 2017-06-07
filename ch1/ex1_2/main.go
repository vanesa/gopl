// Excercise from gopl.io

// ex1_1 prints its command-line arguments.
// Modify the echo program to print the index and value of each of it's arguments, one per line
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var s, sep string

	for i, arg := range os.Args[1:] {
		sep = " "
		s = strconv.Itoa(i) + sep + arg
		fmt.Println(s)
	}
}
