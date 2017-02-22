package main

import (
	"fmt"
	"github.com/jantb/tcell"
	"testing"
)

func TestSyntax(t *testing.T) {
	loadDark()
	syntaxs := syntax([]rune("package main"), "main.go")
	fg, _, _ := syntaxs[0].style.Decompose()
	fmt.Println(fg == tcell.GetColor("#569cd6"))
	fmt.Printf("%x", fg.Hex())
	fmt.Println()
}
func TestSyntax2(t *testing.T) {
	loadDark()
	syntaxs := syntax([]rune("for i, rope := range b.r.Sub(top, length) {"), "main.go")
	fg, _, _ := syntaxs[0].style.Decompose()
	fmt.Println(fg == tcell.GetColor("#569cd6"))
	fmt.Printf("%x", fg.Hex())
	fmt.Println()
}
func TestSyntax3(t *testing.T) {
	loadDark()
	syntaxs := syntax([]rune("type Buffer struct {"), "main.go")
	fg, _, _ := syntaxs[6].style.Decompose()
	fmt.Println(fg == tcell.GetColor("#569cd6"))
	fmt.Printf("%x", fg.Hex())
	fmt.Println()
}
