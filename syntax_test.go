package main

import "testing"

func TestSyntax(t *testing.T) {

	syntax([]rune("func TestSyntax(t *testing.T) {"), "main.go")
}
