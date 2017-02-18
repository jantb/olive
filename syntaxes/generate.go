package main

import (
	"encoding/json"
	"fmt"
	"go/format"
	"io/ioutil"
)

//TmLanguage struct
type TmLanguage struct {
	Comment            string   `json:"comment"`
	FileTypes          []string `json:"fileTypes"`
	FoldingStartMarker string   `json:"foldingStartMarker"`
	FoldingStopMarker  string   `json:"foldingStopMarker"`
	Name               string   `json:"name"`
	Patterns           []struct {
		Begin    string `json:"begin"`
		Captures struct {
			Zero struct {
				Name string `json:"name"`
			} `json:"0"`
		} `json:"captures"`
		Comment string `json:"comment"`
		End     string `json:"end"`
		Name    string `json:"name"`
	} `json:"patterns"`
	Repository struct {
		Brackets struct {
			Patterns []struct {
				Match string `json:"match"`
				Name  string `json:"name"`
			} `json:"patterns"`
		} `json:"brackets"`
		Delimiters struct {
			Patterns []struct {
				Match string `json:"match"`
				Name  string `json:"name"`
			} `json:"patterns"`
		} `json:"delimiters"`
		Keywords struct {
			Patterns []struct {
				Comment string `json:"comment"`
				Match   string `json:"match"`
				Name    string `json:"name"`
			} `json:"patterns"`
		} `json:"keywords"`
		Operators struct {
			Comment  string `json:"comment"`
			Patterns []struct {
				Match string `json:"match"`
				Name  string `json:"name"`
			} `json:"patterns"`
		} `json:"operators"`
		Runes struct {
			Patterns []struct {
				Match string `json:"match"`
				Name  string `json:"name"`
			} `json:"patterns"`
		} `json:"runes"`
		StorageTypes struct {
			Patterns []struct {
				Match string `json:"match"`
				Name  string `json:"name"`
			} `json:"patterns"`
		} `json:"storage_types"`
		StringEscapedChar struct {
			Patterns []struct {
				Match string `json:"match"`
				Name  string `json:"name"`
			} `json:"patterns"`
		} `json:"string_escaped_char"`
		StringPlaceholder struct {
			Patterns []struct {
				Match string `json:"match"`
				Name  string `json:"name"`
			} `json:"patterns"`
		} `json:"string_placeholder"`
		Variables struct {
			Comment  string `json:"comment"`
			Patterns []struct {
				Captures struct {
					One struct {
						Patterns []struct {
							Match string `json:"match"`
							Name  string `json:"name"`
						} `json:"patterns"`
					} `json:"1"`
					Two struct {
						Patterns []struct {
							Include string `json:"include"`
						} `json:"patterns"`
					} `json:"2"`
				} `json:"captures"`
				Match string `json:"match"`
			} `json:"patterns"`
		} `json:"variables"`
	} `json:"repository"`
	ScopeName string `json:"scopeName"`
	Version   string `json:"version"`
}

func main() {
	b, _ := ioutil.ReadFile("go.json")
	t := TmLanguage{}
	json.Unmarshal(b, &t)
	code := fmt.Sprintf("package main\n var golang = %#v", t)
	b, _ = format.Source([]byte(code))
	fmt.Println(string(b))
	ioutil.WriteFile("golang.go", b, 0644)
}
