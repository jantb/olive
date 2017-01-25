package main

import (
	"fmt"

	"github.com/jantb/rope"
)

func main() {
	r := rope.NewFromBytes([]byte("test"))
	fmt.Println(string(r.Bytes()))
	Display()
}
