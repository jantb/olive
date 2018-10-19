package go_plugin

import (
	"bytes"
	"encoding/json"
	"github.com/go-xmlfmt/xmlfmt"
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
	return string(bytes)
}

func FormatXml(content string) string {
	return xmlfmt.FormatXML(content, "", "  ")
}

func FormatJson(content string) string {
	buf := new(bytes.Buffer)
	err := json.Indent(buf, []byte(content), "", "  ")
	if err != nil {
		return content
	}
	return buf.String()
}
