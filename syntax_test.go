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

func TestSyntax4(t *testing.T) {
	loadDark()
	syntaxs := syntax([]rune("make(shds)"), "main.go")
	fg, _, _ := syntaxs[2].style.Decompose()
	fmt.Println(fg == tcell.GetColor("#569cd6"))
	fmt.Printf("%x", fg.Hex())
	fmt.Println()
}
func TestSyntax5(t *testing.T) {
	loadDark()
	syntaxs := syntax([]rune("for i, rope := range b.r.Sub(top, length) {"), "main.go")
	fg, _, _ := syntaxs[2].style.Decompose()
	if fg != tcell.GetColor("#c586c0") {
		fmt.Printf("%x", fg.Hex())
		t.Fail()
	}
	fg, _, _ = syntaxs[16].style.Decompose()
	if fg != tcell.GetColor("#c586c0") {
		fmt.Printf("%x", fg.Hex())
		t.Fail()
	}

}
func TestSyntax6(t *testing.T) {
	loadDark()
	syntaxs := syntax([]rune("file, err := os.Open(filename)"), "main.go")
	fg, _, _ := syntaxs[2].style.Decompose()
	if fg != tcell.GetColor("#9cdcfe") {
		fmt.Printf("%x", fg.Hex())
		t.Fail()
	}

}
