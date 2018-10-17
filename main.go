package main

import (
	"fmt"
	"github.com/jantb/olive/editor"
	"io"
	"log"
	"os"
	"os/exec"
)

type readwriter struct {
	io.Reader
	io.Writer
}

func die(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, format, args...)
	os.Exit(1)
}

func main() {
	configDir := os.Getenv("XI_CONFIG_DIR")
	if configDir == "" {
		configDir = os.Getenv("HOME") + "/.config/xi/"
	}

	path, err := exec.LookPath("xi-core")
	if err != nil {
		die("xi-core was not found in your PATH")
	}

	p := exec.Command(path)
	stdout, _ := p.StdoutPipe()
	stdin, _ := p.StdinPipe()
	if err := p.Start(); err != nil {
		die("error: %v", err.Error())
	}

	f, _ := os.OpenFile("out.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	defer f.Close()
	log.SetOutput(f)

	rw := readwriter{stdout, stdin}
	//editor.NewEditor(rw, configDir)
	//e.Start()
	edit := editor.NewEdit(rw, configDir)
	edit.Start()

}
