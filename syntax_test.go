package main

import (
	"fmt"
	"testing"
)

func TestSyntax6(t *testing.T) {
	transforSyntax()
	for _, s := range syntaxDef {
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
	transforSyntax()
	for _, s := range syntaxDef {
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
	transforSyntax()

	for _, s := range syntaxDef {
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
	transforSyntax()
	for _, s := range syntaxDef {
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
