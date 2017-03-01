package main

import (
	"encoding/json"
	"fmt"
	"go/format"
	"io/ioutil"
	"strings"
)

func main() {
	files, _ := ioutil.ReadDir("./syntaxes")
	syntax := []map[string]interface{}{}
	for _, f := range files {
		if strings.HasSuffix(f.Name(), "json") {
			fmt.Println(f.Name())
			b, _ := ioutil.ReadFile("syntaxes/" + f.Name())
			var dat map[string]interface{}
			json.Unmarshal(b, &dat)
			syntax = append(syntax, dat)
		}
	}
	code := fmt.Sprintf("package main\n var %s = %#v", "syntaxes", syntax)
	b, _ := format.Source([]byte(code))
	ioutil.WriteFile("syntaxes.go", b, 0644)
}
