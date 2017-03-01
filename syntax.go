package main

import (
	"fmt"
	"github.com/dlclark/regexp2"
	"strconv"
	"strings"
)

type Token struct {
	Scope []string
	Loc   []int
}

func syntax(textr []rune, filename string) []Token {
	text := string(textr)
	for _, syntax := range syntaxes {
		filetypes := syntax["fileTypes"]
		for _, language := range filetypes.([]interface{}) {
			if strings.HasSuffix(filename, language.(string)) {
				return highlight(text, syntax)
			}
		}
	}
	return []Token{}
}

var line = []rune{}
var lineOffset = 0

func highlight(lines string, lang map[string]interface{}) []Token {
	scope := []string{}
	tokens := []Token{}
	scope = append(scope, lang["scopeName"].(string))
	for _, l := range strings.Split(lines, "\n") {
		//fmt.Println(line)
		patterns := lang["patterns"]
		repository := lang["repository"]
		line = []rune(l)
		for _, pattern := range patterns.([]interface{}) {
			highlightPattern(scope, &tokens, pattern, repository.(map[string]interface{}))
		}
		fmt.Println(scope)
	}
	return tokens
}

func highlightPattern(scope []string, tokens *[]Token, pattern interface{}, repository map[string]interface{}) {
	p := pattern.(map[string]interface{})
	match := p["match"]
	name := p["name"]
	begin := p["begin"]
	end := p["end"]
	captures := p["captures"]
	beginCaptures := p["beginCaptures"]
	endCaptures := p["endCaptures"]
	patterns := p["patterns"]
	if match != nil {
		r := regexp2.MustCompile(match.(string), 0)
		if m, _ := r.FindRunesMatch(line[lineOffset:]); m != nil {
			handleBeginEnd(name, m, scope, captures, tokens)
		}
		return
	}
	if begin != nil {
		r := regexp2.MustCompile(begin.(string), 0)
		if m, _ := r.FindRunesMatch(line[lineOffset:]); m != nil {
			i := m.GroupByNumber(0).Index
			if i != 0 {
				return
			}
			handleBeginEnd(name, m, scope, beginCaptures, tokens)

			if patterns != nil {
				for _, pattern := range patterns.([]interface{}) {
					highlightPattern(scope, tokens, pattern, repository)
				}
			}
			r := regexp2.MustCompile(end.(string), 0)
			if m, _ := r.FindRunesMatch(line[lineOffset:]); m != nil {
				handleBeginEnd(name, m, scope, endCaptures, tokens)
				fmt.Println("||" + m.String() + "||")
			}
		}
		return
	}

	if patterns != nil {
		for _, pattern := range patterns.([]interface{}) {
			highlightPattern(scope, tokens, pattern, repository)
		}
	}

	include := p["include"]
	if include != nil {
		includeString := include.(string)[1:]
		highlightPattern(scope, tokens, repository[includeString].(map[string]interface{}), repository)
	}

}
func handleBeginEnd(name interface{}, match *regexp2.Match, scope []string, captures interface{}, tokens *[]Token) {
	if name != nil {
		scope = append(scope, name.(string))
	}
	if captures != nil {
		for i, group := range match.Groups() {
			capt := captures.(map[string]interface{})[strconv.Itoa(i)]
			if capt == nil {
				continue
			}
			name := capt.(map[string]interface{})["name"]

			scope = append(scope, name.(string))
			*tokens = append(*tokens, Token{Scope: scope, Loc: []int{group.Index + lineOffset, group.Index + group.Length + lineOffset}})
			lineOffset = lineOffset + match.GroupByNumber(i).Length
		}
	}
}
