package main

import (
	"encoding/json"
	"fmt"
	"go/format"
	"io/ioutil"
	"strings"
)

//TmLanguage struct
type TmLanguage struct {
	ScopeName          string               `json:"scopeName"`
	Name               string               `json:"name"`
	Comment            string               `json:"comment"`
	FileTypes          []string             `json:"fileTypes"`
	FoldingStartMarker string               `json:"foldingStartMarker"`
	FoldingStopMarker  string               `json:"foldingStopMarker"`
	Patterns           []TmLanguagePatterns `json:"patterns"`
	Repository         TmLanguageRepository `json:"repository"`
	Version            string               `json:"version"`
}
type TmLanguagePatterns struct {
	Comment       string              `json:"comment,omitempty"`
	Match         string              `json:"match,omitempty"`
	Captures      TmLanguageCaptures  `json:"captures,omitempty"`
	Name          string              `json:"name,omitempty"`
	Begin         string              `json:"begin,omitempty"`
	BeginCaptures TmLanguageCaptures  `json:"beginCaptures,omitempty"`
	Patterns      []TmLanguagePattern `json:"patterns,omitempty"`
	End           string              `json:"end,omitempty"`
	EndCaptures   TmLanguageCaptures  `json:"endCaptures,omitempty"`
	Include       string              `json:"include,omitempty"`
}
type TmLanguagePattern struct {
	Match    string             `json:"match"`
	Captures TmLanguageCaptures `json:"captures"`
}

type TmLanguageCaptures struct {
	Num0 TmLanguageName `json:"0"`
	Num1 TmLanguageName `json:"1"`
	Num2 TmLanguageName `json:"2"`
	Num3 TmLanguageName `json:"3"`
	Num4 TmLanguageName `json:"4"`
	Num5 TmLanguageName `json:"5"`
}
type TmLanguageName struct {
	Name     string                               `json:"name"`
	Patterns []TmLanguagePatternsNameMatchInclude `json:"patterns,omitempty"`
}

type TmLanguagePatternsNameMatchInclude struct {
	Name    string `json:"name"`
	Match   string `json:"match"`
	Include string `json:"include,omitempty"`
}

type TmLanguageRepository struct {
	Brackets struct {
		Patterns []TmLanguagePatterns `json:"patterns"`
	} `json:"brackets"`
	Delimiters struct {
		Patterns []TmLanguagePatterns `json:"patterns"`
	} `json:"delimiters"`
	Keywords struct {
		Patterns []TmLanguagePatterns `json:"patterns"`
	} `json:"keywords"`
	Operators struct {
		Patterns []TmLanguagePatterns `json:"patterns"`
	} `json:"operators"`
	Runes struct {
		Patterns []TmLanguagePatterns `json:"patterns"`
	} `json:"runes"`
	StorageTypes struct {
		Patterns []TmLanguagePatterns `json:"patterns"`
	} `json:"storage_types"`
	StringEscapedChar struct {
		Patterns []TmLanguagePatterns `json:"patterns"`
	} `json:"string_escaped_char"`
	StringPlaceholder struct {
		Patterns []TmLanguagePatterns `json:"patterns"`
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
}

func main() {
	generate("syntaxes/go.json", "golang")
	generate("syntaxes/json.json", "json")
}
func generate(path, language string) {
	b, _ := ioutil.ReadFile(path)
	t := TmLanguage{}
	json.Unmarshal(b, &t)
	code := fmt.Sprintf("package main\n var %s = %#v", language, t)
	code = strings.Replace(code, "main.", "", -1)
	b, _ = format.Source([]byte(code))
	ioutil.WriteFile(language+".go", b, 0644)
}
