package main

import "testing"

func TestSyntax(t *testing.T) {

	syntax([]rune("package main"), "main.go")
}
