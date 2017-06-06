// Excercise from gopl.io

// ex1_1 prints its command-line arguments.
// Modify the echo program to also print os.Args[0], the name of the command that invoked it.
package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string

	for _, arg := range os.Args[0:] {
		sep = " "
		s += sep + arg
	}

	fmt.Println(s)
}
