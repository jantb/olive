package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/jantb/tcell"
)

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func TermMessage(s tcell.Screen, msg ...interface{}) {
	s.Fini()

	fmt.Println(msg...)
	fmt.Print("\nPress enter to continue")

	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')

}
