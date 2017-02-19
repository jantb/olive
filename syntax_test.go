package main

import (
	"fmt"
	"testing"
)

func TestSyntax(t *testing.T) {
	loadDark()
	syntaxs := syntax([]rune("package main"), "main.go")
	syntaxs = syntax([]rune("package main"), "main.go")
	fmt.Println(syntaxs)
}
