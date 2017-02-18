package main

import "testing"
import "fmt"
import (
	"io/ioutil"
	"strings"
)

func TestBufferInsert(t *testing.T) {
	b := Buffer{}
	b.New()
	text := "Buffer insert test"
	b.Insert(0, 0, []rune(text))
	if text != string(b.GetLines(0, 0, 1, 100)[0]) {
		fmt.Println(text + "!=" + string(b.GetLines(0, 0, 1, 100)[0]))
		t.Fail()
	}
}

func TestBufferInsert2(t *testing.T) {
	b := Buffer{}
	b.New()
	text := "øøø"
	b.Insert(0, 0, []rune(string('ø')))
	b.Insert(0, 0, []rune(string('ø')))
	b.Insert(0, 0, []rune(string('ø')))
	if text != string(b.GetLines(0, 0, 1, 100)[0]) {
		fmt.Println(text + "!=" + string(b.GetLines(0, 0, 1, 100)[0]))
		t.Fail()
	}
}

func TestBufferInsert3(t *testing.T) {
	b := Buffer{}
	b.New()
	text := "øøø"
	b.Insert(0, 0, []rune(string('ø')))
	if "ø" != string(b.GetLines(0, 0, 1, 100)[0]) {
		fmt.Println(text + "!=" + string(b.GetLines(0, 0, 1, 100)[0]))
		t.Fail()
	}
	b.Insert(0, 0, []rune(string('ø')))
	if "øø" != string(b.GetLines(0, 0, 1, 100)[0]) {
		fmt.Println(text + "!=" + string(b.GetLines(0, 0, 1, 100)[0]))
		t.Fail()
	}
	b.Insert(0, 0, []rune(string('ø')))
	if "øøø" != string(b.GetLines(0, 0, 1, 100)[0]) {
		fmt.Println(text + "!=" + string(b.GetLines(0, 0, 1, 100)[0]))
		t.Fail()
	}
}

func TestBufferInsert4(t *testing.T) {
	b := Buffer{}
	b.Open("main.go")
	by, _ := ioutil.ReadFile("main.go")
	if string(by) != string(b.Bytes()) {
		fmt.Println("file content in buffer is not the same as in the file: \"" + string(b.Bytes()) + "\"")
		t.Fail()
	}
}

func TestBufferInsert5(t *testing.T) {
	b := Buffer{}
	b.New()
	expexted := "\n\n\n"
	b.InsertEnter(0, 0)
	b.InsertEnter(1, 0)
	b.InsertEnter(0, 0)
	if expexted != string(b.Bytes()) {
		fmt.Println("\"" + expexted + "\"" + "!=" + "\"" + string(b.Bytes()) + "\"")
		t.Fail()
	}
}

func TestBufferInsert15(t *testing.T) {
	split := strings.Split(string("\n"), "\n")
	fmt.Println(split[1])
	fmt.Println(len(split))
}

func TestBufferInsert7(t *testing.T) {
	b := Buffer{}
	b.New()
	expexted := "hei\n\n"
	b.Insert(0, 0, []rune("hei"))
	b.InsertEnter(0, 3)
	b.InsertEnter(1, 0)
	if expexted != string(b.Bytes()) {
		fmt.Println(expexted + "\n!=\n" + string(b.Bytes()))
		t.Fail()
	}
}

func TestBufferInsert8(t *testing.T) {
	b := Buffer{}
	b.New()
	expexted := "h\ne\ni"
	b.Insert(0, 0, []rune("hei"))
	b.InsertEnter(0, 1)
	b.InsertEnter(1, 1)
	if expexted != string(b.Bytes()) {
		fmt.Println(expexted + "\n!=\n" + string(b.Bytes()))
		t.Fail()
	}
}

func TestBufferInsert9(t *testing.T) {
	b := Buffer{}
	b.New()
	expexted := "h\ne\ni\n"
	b.Insert(0, 0, []rune("hei"))
	b.InsertEnter(0, 1)
	b.InsertEnter(1, 1)
	b.InsertEnter(2, 1)
	if expexted != string(b.Bytes()) {
		fmt.Println(expexted + "\n!=\n" + string(b.Bytes()))
		t.Fail()
	}
}

func TestBufferInsert10(t *testing.T) {
	b := Buffer{}
	b.New()
	expexted := "\n\n"
	b.InsertEnter(0, 0)
	b.InsertEnter(1, 0)
	if expexted != string(b.Bytes()) {
		fmt.Println(expexted + "\n!=\n" + string(b.Bytes()))
		t.Fail()
	}
}

func TestBufferInsert11(t *testing.T) {
	b := Buffer{}
	b.New()
	expexted := ""
	b.InsertEnter(0, 0)
	b.Delete(1, -1)
	if expexted != string(b.Bytes()) {
		fmt.Println(expexted + "\n!=" + string(b.Bytes()))
		t.Fail()
	}
}

func TestBufferInsert12(t *testing.T) {
	b := Buffer{}
	b.New()
	expexted := "\n1"
	b.InsertEnter(0, 0)
	if "\n" != string(b.Bytes()) {
		fmt.Println("\n" + "\n!=" + string(b.Bytes()))
		t.Fail()
	}
	b.Insert(1, 0, []rune("1"))
	if expexted != string(b.Bytes()) {
		fmt.Println("\"" + expexted + "\"" + "\n!=\n" + "\"" + string(b.Bytes()) + "\"")
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
	b.Insert(0, 0, []rune(text))
	b.RemoveRow(0)
	if "" != string(b.GetLines(0, 0, 1, 100)[0]) {
		fmt.Println("!=" + string(b.GetLines(0, 0, 1, 100)[0]))
		t.Fail()
	}
}

func TestBufferDeleteChar(t *testing.T) {
	b := Buffer{}
	b.New()
	text := "Buffer h insert test"
	b.Insert(0, 0, []rune(text))
	b.Delete(0, 7)
	b.Delete(0, 7)
	if "Buffer insert test" != string(b.GetLines(0, 0, 1, 100)[0]) {
		fmt.Println("Buffer insert test" + "!=" + string(b.GetLines(0, 0, 1, 100)[0]))
		t.Fail()
	}
}

func TestBufferBackspace1(t *testing.T) {
	b := Buffer{}
	b.New()
	text := "\n\n\n"
	b.Insert(0, 0, []rune(text))
	b.Backspace(2, 0)
	if "\n\n" != string(b.Bytes()) {
		fmt.Println("\n\n" + "!=" + string(b.Bytes()))
		t.Fail()
	}
}

func TestBufferBackspace2(t *testing.T) {
	b := Buffer{}
	b.New()
	text := "\n\n\n"
	b.Insert(0, 0, []rune(text))
	b.Backspace(1, 0)
	b.Backspace(1, 0)
	if "\n" != string(b.Bytes()) {
		fmt.Println("\n" + "!=" + string(b.Bytes()))
		t.Fail()
	}
}

func TestBufferBackspace3(t *testing.T) {
	b := Buffer{}
	b.New()
	text := "\n\n\nHello"
	b.Insert(0, 0, []rune(text))
	b.Backspace(1, 0)
	if "\n\nHello" != string(b.Bytes()) {
		fmt.Println("\n\nHello" + "!=" + string(b.Bytes()))
		t.Fail()
	}
	b.Backspace(1, 0)
	if "\nHello" != string(b.Bytes()) {
		fmt.Println("\"\nHello" + "!=" + string(b.Bytes()) + "\"")
		t.Fail()
	}
}

func TestBufferBackspace4(t *testing.T) {
	b := Buffer{}
	b.New()
	text := "Hello"
	b.Insert(0, 0, []rune(text))
	b.Backspace(0, 0)
	if "Hello" != string(b.Bytes()) {
		fmt.Println("Hello" + "!=" + string(b.Bytes()))
		t.Fail()
	}
}

func TestBufferBackspace5(t *testing.T) {
	b := Buffer{}
	b.New()
	text := "Hello"
	b.Insert(0, 0, []rune(text))
	b.Backspace(1, 0)
	if "Hello" != string(b.Bytes()) {
		fmt.Println("Hello" + "!=" + string(b.Bytes()))
		t.Fail()
	}
}
