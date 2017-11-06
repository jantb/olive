package main

import "testing"
import "fmt"
import "github.com/stretchr/testify/assert"
import (
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

func TestBufferInsert_1(t *testing.T) {
	b := Buffer{}
	b.NewText("Buffer insert test")
	text := "Buffer insert test"
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

func TestBufferInsert71(t *testing.T) {
	b := Buffer{}
	b.New()
	expexted := "ieh"
	b.Insert(0, 0, []rune("h"))
	b.Insert(0, 0, []rune("e"))
	b.Insert(0, 0, []rune("i"))
	if expexted != string(b.Bytes()) {
		fmt.Println(expexted + "\n!=\n" + string(b.Bytes()))
		t.Fail()
	}
}

func TestBufferInsert72(t *testing.T) {
	b := Buffer{}
	b.New()
	expexted := "hei"
	b.Insert(0, 0, []rune("h"))
	b.Insert(0, 1, []rune("e"))
	b.Insert(0, 2, []rune("i"))
	if expexted != string(b.Bytes()) {
		fmt.Println(expexted + "\n!=\n" + string(b.Bytes()))
		t.Fail()
	}
}

func TestBufferInsert74(t *testing.T) {
	b := Buffer{}
	b.New()
	expexted := "\n1"
	b.InsertEnter(0, 0)
	b.Insert(1, 0, []rune("1"))
	if expexted != string(b.Bytes()) {
		fmt.Println(expexted + "\n!=\n" + string(b.Bytes()))
		t.Fail()
	}
}

func TestBufferInsert73(t *testing.T) {
	b := Buffer{}
	b.New()
	expexted := "hei\nhei"
	b.Insert(0, 0, []rune("h"))
	b.Insert(0, 1, []rune("e"))
	b.Insert(0, 2, []rune("i"))

	b.InsertEnter(0, 3)
	if "hei\n" != string(b.Bytes()) {
		fmt.Println("hei\n" + "\n!=\n" + string(b.Bytes()))
		t.Fail()
	}
	b.Insert(1, 0, []rune("h"))
	b.Insert(1, 1, []rune("e"))
	b.Insert(1, 2, []rune("i"))
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
	b.Delete(1, 0)
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

func TestBufferGetLineLength(t *testing.T) {
	b := Buffer{}
	b.New()
	text := "123456\n\nabcdef\n98765"
	b.Insert(0, 0, []rune(text))
	len1 := b.GetLineLen(0)
	len2 := b.GetLineLen(1)
	len3 := b.GetLineLen(2)
	assert.Equal(t, 6, len1, "")
	assert.Equal(t, 0, len2, "")
	assert.Equal(t, 6, len3, "")
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

func TestGetLines2(t *testing.T) {
	b := Buffer{}
	b.New()
	text := "1\n2\n3\n4"
	b.Insert(0, 0, []rune(text))
	if text != string(b.Bytes()) {
		fmt.Println(text + "!=" + string(b.Bytes()))
		t.Fail()
	}
	line := b.GetLine(1)
	if "2" != string(line) {
		fmt.Println("2" + "!=" + string(line))
		t.Fail()
	}
}

func TestGetLines(t *testing.T) {
	b := Buffer{}
	b.NewText(`package main

import (
	"os"
	"path/filepath"
	"syscall"
	"log"
	"os/user"
)

var buffer = Buffer{}

func main() {

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	logFile, _ := os.OpenFile(filepath.Join(usr.HomeDir, ".olive.log"), os.O_WRONLY|os.O_CREATE|os.O_SYNC, 0755)
	syscall.Dup2(int(logFile.Fd()), 1)
	syscall.Dup2(int(logFile.Fd()), 2)

	if len(os.Args) == 2 {
		buffer.Open(os.Args[1])
		loadDark()
		transforSyntax()
		Display(&buffer)
	}
}
`)
	lines := b.GetLines(0, 0, 59, 100)
	for _, value := range lines {
		fmt.Println(string(value))
	}
}
