package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/jantb/tcell"
)

// Min of int
func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// Max of int
func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// TermMessage Quit and sgow termmessage
func TermMessage(s tcell.Screen, msg ...interface{}) {
	s.Fini()

	fmt.Println(msg...)
	fmt.Print("\nPress enter to continue")

	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')

}
