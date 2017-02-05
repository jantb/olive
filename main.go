package main

import (
	"os"
)

var buffer = Buffer{}

func main() {
	if len(os.Args) == 2 {
		buffer.Open(os.Args[1])
		Display(&buffer)
	}
}
