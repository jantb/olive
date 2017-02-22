package main

import (
	"regexp"
	"strings"

	"github.com/jantb/tcell"
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
	Name string `json:"name"`
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

func syntax(liner []rune, filename string) []Backing {
	line := []byte(string(liner))
	backing := make([]Backing, len(liner), len(liner))
	for key := range backing {
		backing[key].style = getEditorStyle()
	}
	if strings.HasSuffix(filename, ".go") {
		for _, pattern := range golang.Patterns {
			// reBegin := regexp.MustCompile(pattern.Begin)
			// reEnd := regexp.MustCompile(pattern.End)
			// loc := reBegin.FindIndex(line)
			// loc = reEnd.FindIndex(line)
			if pattern.Include != "" {
				switch pattern.Include {
				case "#brackets":
					locs, styles := getStyle(golang.Repository.Keywords.Patterns, line)
					backing, liner = style(locs, backing, styles, liner)
				case "#delimiters":
					locs, styles := getStyle(golang.Repository.Keywords.Patterns, line)
					backing, liner = style(locs, backing, styles, liner)
				case "#keywords":
					locs, styles := getStyle(golang.Repository.Keywords.Patterns, line)
					backing, liner = style(locs, backing, styles, liner)
				case "#operators":
					locs, styles := getStyle(golang.Repository.Keywords.Patterns, line)
					backing, liner = style(locs, backing, styles, liner)
				case "#runes":
					locs, styles := getStyle(golang.Repository.Keywords.Patterns, line)
					backing, liner = style(locs, backing, styles, liner)
				case "#storage_types":
					locs, styles := getStyle(golang.Repository.Keywords.Patterns, line)
					backing, liner = style(locs, backing, styles, liner)
				case "#string_escaped_char":
					locs, styles := getStyle(golang.Repository.Keywords.Patterns, line)
					backing, liner = style(locs, backing, styles, liner)
				case "#string_placeholder":
					locs, styles := getStyle(golang.Repository.Keywords.Patterns, line)
					backing, liner = style(locs, backing, styles, liner)
				}

			}
			if pattern.Match != "" {
				reMatch := regexp.MustCompile(pattern.Match)
				locs := reMatch.FindAllIndex(line, -1)
				if locs == nil {
					continue
				}
				l := len(locs)
				if l > 0 {
					backing = colorLock(locs[0], backing, liner, pattern.Captures.Num1.Name)
				}
				if l > 1 {
					backing = colorLock(locs[1], backing, liner, pattern.Captures.Num2.Name)
				}
				if l > 2 {
					backing = colorLock(locs[2], backing, liner, pattern.Captures.Num3.Name)
				}
				if l > 3 {
					backing = colorLock(locs[3], backing, liner, pattern.Captures.Num4.Name)
				}
				if l > 4 {
					backing = colorLock(locs[4], backing, liner, pattern.Captures.Num5.Name)
				}

			}
		}
	}
	return backing
}
func colorLock(loc []int, backing []Backing, liner []rune, name string) []Backing {
	if loc != nil && name != "" {
		style := getThemeColor(name)
		for index := loc[0]; index < loc[1]; index++ {
			b := backing[index]
			b.style = style
			b.value = liner[index]
			backing[index] = b
		}
	}
	return backing
}

func style(locs [][]int, backing []Backing, styles []tcell.Style, liner []rune) ([]Backing, []rune) {
	for key, loc := range locs {
		for index := loc[0]; index < loc[1]; index++ {
			b := backing[index]
			b.style = styles[key]
			b.value = liner[index]
			backing[index] = b
		}
	}
	return backing, liner
}

func getStyle(patterns []TmLanguagePatterns, line []byte) ([][]int, []tcell.Style) {
	locs := [][]int{}
	styles := []tcell.Style{}
	for _, pattern := range patterns {
		if pattern.Match != "" {
			reMatch := regexp.MustCompile(pattern.Match)
			loc := reMatch.FindIndex(line)
			if loc != nil {
				styles = append(styles, getThemeColor(pattern.Name))
				locs = append(locs, loc)
			}
		}
	}
	return locs, styles
}
