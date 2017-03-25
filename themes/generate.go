package main

import (
	"encoding/json"
	"fmt"
	"go/format"
	"io/ioutil"
	"strings"
)

//go:generate go run generate.go
type Theme struct {
	Name     string `json:"name"`
	Settings []struct {
		Scope    string   `json:"scope"`
		Scopes   []string `json:"scopes"`
		Settings struct {
			FontStyle  string `json:"fontStyle,omitempty"`
			Foreground string `json:"foreground,omitempty"`
			Background string `json:"background,omitempty"`
		} `json:"settings"`
		Name string `json:"name,omitempty"`
	} `json:"settings"`
}

func main() {
	b, _ := ioutil.ReadFile("dark-vs.json")
	t := Theme{}
	json.Unmarshal(b, &t)
	code := fmt.Sprintf("package main\n var dark = %#v", t)
	code = strings.Replace(code, "main.", "", -1)
	b, _ = format.Source([]byte(code))
	ioutil.WriteFile("../themes.go", b, 0644)
}
