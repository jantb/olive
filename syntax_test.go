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
	syntax := transforSyntax()
	for _, s := range syntax {
		for _, f := range s.FileTypes {
			if f == "go" {
				tokens := highlightLine([]rune("//Loc holds a location"), s)
				for _, token := range tokens {
					fmt.Println(token)
				}
				return
			}
		}
	}
}
func TestSyntax7(t *testing.T) {
	syntax := transforSyntax()
	for _, s := range syntax {
		for _, f := range s.FileTypes {
			if f == "json" {
				tokens := highlightLine([]rune("{}"), s)
				for _, token := range tokens {
					fmt.Println(token)
				}
				return
			}
		}
	}
}

func TestSyntax8(t *testing.T) {
	syntax := transforSyntax()
	for _, s := range syntax {
		for _, f := range s.FileTypes {
			if f == "json" {
				tokens := highlightLine([]rune("{\"h\":\"x\"}"), s)
				for _, token := range tokens {
					fmt.Println(token)
				}
				return
			}
		}
	}
}
func TestSyntax9(t *testing.T) {
	syntax := transforSyntax()
	for _, s := range syntax {
		for _, f := range s.FileTypes {
			if f == "go" {
				tokens := highlightLine([]rune("package main"), s)
				for _, token := range tokens {
					fmt.Println(token)
				}
				return
			}
		}
	}
}
