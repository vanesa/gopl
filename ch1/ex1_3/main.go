// Excercise from gopl.io

// ex1_3 benchmarks the echo functions.
// Experiment to measure the difference in running time between our potentially
// inefficient version and the one that uses strings.Join.
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	echo1()
	echo2()
	echo3()
}

func echo1() {
	start := time.Now()
	var s, sep string

	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}

	secs := time.Since(start).Seconds()
	fmt.Printf("%.12fs, %s\n", secs, "echo1")
}

func echo2() {
	start := time.Now()
	var s, sep string

	for _, arg := range os.Args[1:] {
		sep = " "
		s += sep + arg
	}

	secs := time.Since(start).Seconds()
	fmt.Printf("%.12fs, %s\n", secs, "echo2")
}

func echo3() {
	start := time.Now()
	secs := time.Since(start).Seconds()
	fmt.Printf("%.12fs, %s, %s\n", secs, strings.Join(os.Args[1:], " "), "echo2")
}
