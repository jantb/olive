package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/dlclark/regexp2"
)

type Token struct {
	Scope []string
	Loc   []int
}

type Syntax struct {
	ScopeName      string
	Name           string
	FileTypes      []string
	FirstLineMatch regexp2.Regexp
	Patterns       []Pattern
	Repository     map[string]Pattern
}

type Pattern struct {
	Comment       string
	Match         regexp2.Regexp
	Captures      map[int]Pattern
	Begin         regexp2.Regexp
	BeginCaptures map[int]Pattern
	End           regexp2.Regexp
	EndCaptures   map[int]Pattern
	Name          string
	Patterns      []Pattern
	Include       string
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

func highlightLine(line []rune, syntax Syntax) []Token {
	tokens := []Token{}
	scope := []string{syntax.ScopeName}
	for _, pattern := range syntax.Patterns {
		highlightLinePattern(syntax, &tokens, scope, line, pattern)
	}
	lineOffset = 0
	return tokens
}

func highlightLinePattern(syntax Syntax, tokens *[]Token, scope []string, line []rune, pattern Pattern) {
	if len(line) < lineOffset {
		return
	}
	if pattern.Match.String() != "" {
		if m, _ := pattern.Match.FindRunesMatch(line[lineOffset:]); m != nil {
			if pattern.Name != "" {
				scope = append(scope, pattern.Name)
			}
			handleMatch(m, scope, pattern.Captures, tokens)
		}
	}
	if pattern.Begin.String() != "" {
		if m, _ := pattern.Begin.FindRunesMatch(line[lineOffset:]); m != nil {
			if pattern.Name != "" {
				scope = append(scope, pattern.Name)
			}
			if len(pattern.Captures) > 0 {
				handleMatch(m, scope, pattern.Captures, tokens)
			} else {
				handleMatch(m, scope, pattern.BeginCaptures, tokens)
			}
			for _, p := range pattern.Patterns {
				highlightLinePattern(syntax, tokens, scope, line, p)
			}
			if m, _ := pattern.End.FindRunesMatch(line[lineOffset:]); m != nil {
				handleMatch(m, scope, pattern.EndCaptures, tokens)
			}
		}
		return
	}
	for _, p := range pattern.Patterns {
		highlightLinePattern(syntax, tokens, scope, line, p)
	}
	if pattern.Include != "" {
		highlightLinePattern(syntax, tokens, scope, line, syntax.Repository[pattern.Include[1:]])
	}

}

func transforSyntax() []Syntax {
	syntaxDef := []Syntax{}
	for _, syntax := range syntaxes {
		syn := Syntax{
			ScopeName: syntax["scopeName"].(string),
			Name:      syntax["name"].(string),
		}
		syn.Repository = make(map[string]Pattern)
		if syntax["firstLineMatch"] != nil {
			syn.FirstLineMatch = *regexp2.MustCompile(syntax["firstLineMatch"].(string), 0)
		}

		filetypes := []string{}
		for _, filetype := range syntax["fileTypes"].([]interface{}) {
			filetypes = append(filetypes, filetype.(string))
		}
		syn.FileTypes = filetypes

		for _, pattern := range syntax["patterns"].([]interface{}) {
			syn.Patterns = append(syn.Patterns, transferPattern(pattern))
		}

		if syntax["repository"] != nil {
			for key, pattern := range syntax["repository"].(map[string]interface{}) {
				syn.Repository[key] = transferPattern(pattern)
			}
		}
		syntaxDef = append(syntaxDef, syn)
	}
	return syntaxDef
}
func transferPattern(pattern interface{}) Pattern {
	p := pattern.(map[string]interface{})
	pat := Pattern{}
	pat.Captures = make(map[int]Pattern)
	pat.BeginCaptures = make(map[int]Pattern)
	pat.EndCaptures = make(map[int]Pattern)
	if p["comment"] != nil {
		pat.Comment = p["comment"].(string)
	}
	if p["name"] != nil {
		pat.Name = p["name"].(string)
	}
	if p["include"] != nil {
		pat.Include = p["include"].(string)
	}

	if p["match"] != nil {
		pat.Match = *regexp2.MustCompile(p["match"].(string), 0)
	}
	if p["begin"] != nil {
		pat.Begin = *regexp2.MustCompile(p["begin"].(string), 0)
	}
	if p["end"] != nil {
		pat.End = *regexp2.MustCompile(p["end"].(string), 0)
	}

	if p["captures"] != nil {
		for key, pattern := range p["captures"].(map[string]interface{}) {
			atoi, _ := strconv.Atoi(key)
			pat.Captures[atoi] = transferPattern(pattern)
		}
	}
	if p["beginCaptures"] != nil {
		for key, pattern := range p["beginCaptures"].(map[string]interface{}) {
			atoi, _ := strconv.Atoi(key)
			pat.BeginCaptures[atoi] = transferPattern(pattern)
		}
	}
	if p["endCaptures"] != nil {
		for key, pattern := range p["endCaptures"].(map[string]interface{}) {
			atoi, _ := strconv.Atoi(key)
			pat.EndCaptures[atoi] = transferPattern(pattern)
		}
	}
	if p["patterns"] != nil {
		for _, pattern := range p["patterns"].([]interface{}) {
			pat.Patterns = append(pat.Patterns, transferPattern(pattern))
		}
	}
	return pat
}

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

func handleMatch(match *regexp2.Match, scope []string, captures map[int]Pattern, tokens *[]Token) {
	if len(captures) > 0 {
		for i, group := range match.Groups() {
			pattern := captures[i]
			scope = append(scope, pattern.Name)
			*tokens = append(*tokens, Token{Scope: scope, Loc: []int{group.Index + lineOffset, group.Index + group.Length + lineOffset}})
			lineOffset = lineOffset + match.GroupByNumber(i).Length
		}
	} else {
		*tokens = append(*tokens, Token{Scope: scope, Loc: []int{lineOffset, match.Index + lineOffset}})
		lineOffset = lineOffset + match.Index
	}
}
