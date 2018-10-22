package go_plugin

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
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

type xmlany struct {
	XMLName xml.Name
	Attrs   []xml.Attr `xml:",any,attr"`
	Content string     `xml:",chardata"`
	Comment string     `xml:",comment"`
	Nested  []*xmlany  `xml:",any"`
}

func FormatXml(content string) string {
	// go get golang.org/x/tools/cmd/goimports
	path, err := exec.LookPath("xmllint")
	if err != nil {
		log.Println("glibxml2-utilswas not found in your PATH")
		return content
	}

	p := exec.Command(path, "--format", "-")
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

func FormatJson(content string) string {
	buf := new(bytes.Buffer)
	err := json.Indent(buf, []byte(content), "", "  ")
	if err != nil {
		return content
	}
	return buf.String()
}
