package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"

	"github.com/jantb/olive/editor"
)
import _ "net/http/pprof"
import _ "net/http"

type readwriter struct {
	io.Reader
	io.Writer
}

func die(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, format, args...)
	os.Exit(1)
}
func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
	configDir := os.Getenv("XI_CONFIG_DIR")
	if configDir == "" {
		configDir = os.Getenv("HOME") + "/.config/xi/"
	}

	path, err := exec.LookPath("xi-core")
	if err != nil {
		die("xi-core was not found in your PATH\n" +
			"Install it from https://github.com/xi-editor/xi-editor/")
	}

	p := exec.Command(path)
	stdout, _ := p.StdoutPipe()
	stdin, _ := p.StdinPipe()
	if err := p.Start(); err != nil {
		die("error: %v", err.Error())
	}

	f, _ := os.OpenFile(os.Getenv("HOME")+"/.olive.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	defer f.Close()
	log.SetOutput(f)

	rw := readwriter{stdout, stdin}
	editor := editor.NewEdit(rw, configDir)
	editor.Start()
	editor.Quit()
}
