package main

import "testing"
import "fmt"
import "io/ioutil"

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
		fmt.Println("file content in buffer is not the same as in the file: " + string(b.Bytes()))
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
		fmt.Println(expexted + "!=" + string(b.Bytes()))
		t.Fail()
	}
}

func TestBufferInsert6(t *testing.T) {
	b := Buffer{}
	b.New()
	expexted := "hei\npå\ndeg\n"
	b.Insert(0, 0, []rune("heipådeg"))
	b.InsertEnter(0, 3)
	b.InsertEnter(1, 2)
	b.InsertEnter(3, 3)
	if expexted != string(b.Bytes()) {
		fmt.Println(expexted + "\n!=\n" + string(b.Bytes()))
		t.Fail()
	}
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
