package main

import "testing"
import "fmt"

func TestBufferInsert(t *testing.T) {
	b := Buffer{}
	b.New()
	text := "Buffer insert test"
	b.Insert(0, 0, []byte(text))
	if text+"\n" != string(b.GetLines(0, 1)[0]) {
		fmt.Println(text + "!=" + string(b.GetLines(0, 1)[0]))
		t.Fail()
	}
}
