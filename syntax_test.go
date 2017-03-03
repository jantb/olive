package main

import (
	"fmt"
	"testing"
)

func TestSyntaxCurlyStartJson(t *testing.T) {
	loadDark()
	syntax := syntax([]rune("{}"), "main.json")
	for _, s := range syntax {
		fmt.Println(s)
	}
}

func TestSyntaxCurlyStartJson2(t *testing.T) {
	loadDark()
	syntax := syntax([]rune("{\"h\":\"x\"}"), "main.json")
	for _, s := range syntax {
		fmt.Println(s)
	}
}

func TestSyntax3(t *testing.T) {
	loadDark()
	syntax := syntax([]rune("package main"), "main.go")
	for _, s := range syntax {
		fmt.Println(s)
	}
}

func TestSyntax4(t *testing.T) {
	loadDark()
	syntax := syntax([]rune("//Loc holds a location"), "main.go")
	for _, s := range syntax {
		fmt.Println(s)
	}
}
func TestSyntax5(t *testing.T) {
	loadDark()
	syntax := syntax([]rune("<h1></h1>"), "main.xml")
	for _, s := range syntax {
		fmt.Println(s)
	}
}

func TestSyntax6(t *testing.T) {
	transforSyntax()
}
