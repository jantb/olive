package main

import "testing"
import "fmt"
import "io/ioutil"

func TestBufferInsert(t *testing.T) {
	b := Buffer{}
	b.New()
	text := "Buffer insert test"
	b.Insert(0, 0, []byte(text))
	if text != string(b.GetLines(0, 1, 100)[0]) {
		fmt.Println(text + "!=" + string(b.GetLines(0, 1, 100)[0]))
		t.Fail()
	}
}

func TestBufferInsert2(t *testing.T) {
	b := Buffer{}
	b.New()
	text := "øøø"
	b.Insert(0, 0, []byte(string('ø')))
	b.Insert(0, 0, []byte(string('ø')))
	b.Insert(0, 0, []byte(string('ø')))
	if text != string(b.GetLines(0, 1, 100)[0]) {
		fmt.Println(text + "!=" + string(b.GetLines(0, 1, 100)[0]))
		t.Fail()
	}
}

func TestBufferInsert3(t *testing.T) {
	b := Buffer{}
	b.Open("main.go")
	by, _ := ioutil.ReadFile("main.go")
	if string(by) != string(b.Bytes()) {
		fmt.Println("file content in buffer is not the same as in the file: " + string(b.Bytes()))
		t.Fail()
	}
}

func TestBufferDeleteEmpty(t *testing.T) {
	b := Buffer{}
	b.New()
	b.RemoveRow(0)
}

func TestBufferDelete(t *testing.T) {
	b := Buffer{}
	b.New()
	text := "Buffer insert test"
	b.Insert(0, 0, []byte(text))
	b.RemoveRow(0)
	if "" != string(b.GetLines(0, 1, 100)[0]) {
		fmt.Println("!=" + string(b.GetLines(0, 1, 100)[0]))
		t.Fail()
	}
}

func TestBufferDeleteChar(t *testing.T) {
	b := Buffer{}
	b.New()
	text := "Buffer h insert test"
	b.Insert(0, 0, []byte(text))
	b.Delete(0, 7)
	b.Delete(0, 7)
	if "Buffer insert test" != string(b.GetLines(0, 1, 100)[0]) {
		fmt.Println("Buffer insert test" + "!=" + string(b.GetLines(0, 1, 100)[0]))
		t.Fail()
	}
}
