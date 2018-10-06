package goPlugin

import (
	"io"
	"log"
	"os/exec"
)

func Format(content string) string {
	// go get golang.org/x/tools/cmd/goimports
	path, err := exec.LookPath("goimports")
	if err != nil {
		log.Println("goimports was not found in your PATH go get golang.org/x/tools/cmd/goimports")
		return content
	}

	p := exec.Command(path)
	stdin, _ := p.StdinPipe()

	io.WriteString(stdin, content)
	stdin.Close()
	bytes, err := p.CombinedOutput()
	if err != nil {
		log.Println(err)
		return content
	}
	log.Println("ok")
	return string(bytes)
}
