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

//func TestSyntaxCurlyStartJson2(t *testing.T) {
//	loadDark()
//	syntax := syntax([]rune("{\"h\":\"x\"}"), "main.json")
//	for _, s := range syntax {
//		fmt.Println(s)
//	}
//}
