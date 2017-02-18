package main

import (
	"testing"
)

func TestCursor_MoveDown(t *testing.T) {
	buffer = Buffer{}
	topRow = 0
	height = 100
	width = 100
	buffer.New()
	text := "Hello\n\n"
	c := GetCursor()
	buffer.Insert(0, 0, []rune(text))
	c.MoveDown()
	if c.loc.row != 1 {
		t.Log(c.loc)
		t.Fail()
	}
	c.MoveDown()
	if c.loc.row != 2 {
		t.Log(c.loc)
		t.Fail()
	}
}
