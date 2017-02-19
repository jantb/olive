package main

import (
	"fmt"
	"testing"
)

func TestSyntax(t *testing.T) {

	syntax := syntax([]rune("package main"), "main.go")
	fmt.Println(syntax)
}
