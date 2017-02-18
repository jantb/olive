package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

//TmLanguage struct
type TmLanguage struct {
	ScopeName          string   `json:"scopeName"`
	Name               string   `json:"name"`
	Comment            string   `json:"comment"`
	FileTypes          []string `json:"fileTypes"`
	FoldingStartMarker string   `json:"foldingStartMarker"`
	FoldingStopMarker  string   `json:"foldingStopMarker"`
	Patterns           []struct {
		Comment  string `json:"comment,omitempty"`
		Match    string `json:"match,omitempty"`
		Captures struct {
			Num1 struct {
				Name string `json:"name"`
			} `json:"1"`
		} `json:"captures,omitempty"`
		Name          string `json:"name,omitempty"`
		Begin         string `json:"begin,omitempty"`
		BeginCaptures struct {
			Num1 struct {
				Name string `json:"name"`
			} `json:"1"`
		} `json:"beginCaptures,omitempty"`
		Patterns []struct {
			Match    string `json:"match"`
			Captures struct {
				Num1 struct {
					Name string `json:"name"`
				} `json:"1"`
				Num2 struct {
					Name string `json:"name"`
				} `json:"2"`
				Num3 struct {
					Name string `json:"name"`
				} `json:"3"`
				Num4 struct {
					Name string `json:"name"`
				} `json:"4"`
				Num5 struct {
					Name string `json:"name"`
				} `json:"5"`
			} `json:"captures"`
		} `json:"patterns,omitempty"`
		End         string `json:"end,omitempty"`
		EndCaptures struct {
			Num0 struct {
				Name string `json:"name"`
			} `json:"0"`
		} `json:"endCaptures,omitempty"`
		Include string `json:"include,omitempty"`
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
				Comment string `json:"comment,omitempty"`
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
				Match    string `json:"match"`
				Captures struct {
					Num1 struct {
						Patterns []struct {
							Match   string `json:"match,omitempty"`
							Name    string `json:"name,omitempty"`
							Include string `json:"include,omitempty"`
						} `json:"patterns"`
					} `json:"1"`
					Num2 struct {
						Patterns []struct {
							Include string `json:"include"`
						} `json:"patterns"`
					} `json:"2"`
				} `json:"captures"`
			} `json:"patterns"`
		} `json:"variables"`
	} `json:"repository"`
	Version string `json:"version"`
}

func syntax(liner []rune, filename string) {
	line := []byte(string(liner))
	b, _ := ioutil.ReadFile("syntaxes/go.json")
	t := TmLanguage{}
	json.Unmarshal(b, &t)
	if strings.HasSuffix(filename, ".go") {
		for _, pattern := range t.Patterns {
			// reBegin := regexp.MustCompile(pattern.Begin)
			// reEnd := regexp.MustCompile(pattern.End)
			// loc := reBegin.FindIndex(line)
			// loc = reEnd.FindIndex(line)
			if pattern.Match != "" {
				reMatch := regexp.MustCompile(pattern.Match)
				loc := reMatch.FindIndex(line)
				if loc != nil {
					l := string(line)
					fmt.Println(pattern.Name)
					fmt.Println(l)
				}
			}
		}
	}
}
