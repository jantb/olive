package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func syntax(textr []rune, filename string) []string {
	text := string(textr)
	//for key := range backing {
	//	backing[key].style = getEditorStyle()
	//}
	for _, syntax := range syntaxes {
		filetypes := syntax["fileTypes"]
		for _, language := range filetypes.([]interface{}) {
			if strings.HasSuffix(filename, language.(string)) {
				return highlight(text, syntax)
			}
		}
	}
	return []string{}
}

func highlight(lines string, lang map[string]interface{}) []string {
	scope := []string{}
	scope = append(scope, lang["scopeName"].(string))
	for _, line := range strings.Split(lines, "\n") {
		//fmt.Println(line)
		patterns := lang["patterns"]
		repository := lang["repository"]
		for _, pattern := range patterns.([]interface{}) {
			highlightPattern(&scope, []byte(line), pattern, repository.(map[string]interface{}))
		}
		fmt.Println(scope)
		break
	}
	return scope
}

func highlightPattern(scope *[]string, line []byte, pattern interface{}, repository map[string]interface{}) {
	p := pattern.(map[string]interface{})
	match := p["match"]
	name := p["name"]
	begin := p["begin"]
	beginCaptures := p["beginCaptures"]
	patterns := p["patterns"]
	if match != nil {
		r := regexp.MustCompile(match.(string))
		locs := r.FindAllSubmatchIndex(line, -1)
		if name != nil && locs != nil {
			fmt.Println(name)
			fmt.Println(string(line))
			fmt.Println(locs)
			*scope = append(*scope, name.(string))
			fmt.Println(scope)
		}
		return
	}
	if begin != nil {
		r := regexp.MustCompile(begin.(string))
		locs := r.FindAllSubmatchIndex(line, -1)
		if name != nil && locs != nil {
			fmt.Println(name)
			fmt.Println(string(line))
			fmt.Println(locs)
			*scope = append(*scope, name.(string))
			fmt.Println(scope)
		}
		if beginCaptures != nil {
			for i, _ := range locs {
				capt := beginCaptures.(map[string]interface{})[strconv.Itoa(i)]
				name := capt.(map[string]interface{})["name"]
				*scope = append(*scope, name.(string))
				fmt.Println(scope)
			}
		}

		return
	}
	if patterns != nil {
		for _, pattern := range patterns.([]interface{}) {
			highlightPattern(scope, []byte(line), pattern, repository)
		}
	}

	include := p["include"]
	if include != nil {
		includeString := include.(string)[1:]
		fmt.Println(includeString)
		highlightPattern(scope, []byte(line), repository[includeString].(map[string]interface{}), repository)
	}

}

//func highlight(line []byte, backing []Backing, liner []rune, lang map[string]interface{}) []Backing {
//
//	for _, pattern := range lang.Patterns {
//		// reBegin := regexp.MustCompile(pattern.Begin)
//		// reEnd := regexp.MustCompile(pattern.End)
//		// loc := reBegin.FindIndex(line)
//		// loc = reEnd.FindIndex(line)
//		if pattern.Include != "" {
//			switch pattern.Include {
//			case "#brackets":
//				locs, styles := getStyle(golang.Repository.Keywords.Patterns, line)
//				backing, liner = style(locs, backing, styles, liner)
//			case "#delimiters":
//				locs, styles := getStyle(golang.Repository.Keywords.Patterns, line)
//				backing, liner = style(locs, backing, styles, liner)
//			case "#keywords":
//				locs, styles := getStyle(golang.Repository.Keywords.Patterns, line)
//				backing, liner = style(locs, backing, styles, liner)
//			case "#operators":
//				locs, styles := getStyle(golang.Repository.Keywords.Patterns, line)
//				backing, liner = style(locs, backing, styles, liner)
//			case "#runes":
//				locs, styles := getStyle(golang.Repository.Keywords.Patterns, line)
//				backing, liner = style(locs, backing, styles, liner)
//			case "#storage_types":
//				locs, styles := getStyle(golang.Repository.Keywords.Patterns, line)
//				backing, liner = style(locs, backing, styles, liner)
//			case "#string_escaped_char":
//				locs, styles := getStyle(golang.Repository.Keywords.Patterns, line)
//				backing, liner = style(locs, backing, styles, liner)
//			case "#string_placeholder":
//				locs, styles := getStyle(golang.Repository.Keywords.Patterns, line)
//				backing, liner = style(locs, backing, styles, liner)
//			}
//
//		}
//		if pattern.Match != "" {
//			reMatch := regexp.MustCompile(pattern.Match)
//			locs := reMatch.FindAllSubmatchIndex(line, -1)
//			if locs == nil {
//				continue
//			}
//			for _, ll := range locs {
//
//				l := len(ll)
//				if l > 3 {
//					backing = colorLock([]int{ll[2], ll[3]}, backing, liner, pattern.Captures.Num1.Name)
//				}
//				if l > 5 {
//					backing = colorLock([]int{ll[4], ll[5]}, backing, liner, pattern.Captures.Num2.Name)
//				}
//				if l > 7 {
//					backing = colorLock([]int{ll[6], ll[7]}, backing, liner, pattern.Captures.Num3.Name)
//				}
//				if l > 9 {
//					backing = colorLock([]int{ll[8], ll[9]}, backing, liner, pattern.Captures.Num4.Name)
//				}
//				if l > 11 {
//					backing = colorLock([]int{ll[10], ll[11]}, backing, liner, pattern.Captures.Num5.Name)
//				}
//			}
//
//		}
//	}
//	return backing
//}
//func colorLock(loc []int, backing []Backing, liner []rune, name string) []Backing {
//	if loc != nil && name != "" {
//		style := getThemeColor(name)
//		for index := loc[0]; index < loc[1]; index++ {
//			b := backing[index]
//			b.style = style
//			b.value = liner[index]
//			backing[index] = b
//		}
//	}
//	return backing
//}
//
//func style(locs [][]int, backing []Backing, styles []tcell.Style, liner []rune) ([]Backing, []rune) {
//	for key, loc := range locs {
//		for index := loc[0]; index < loc[1]; index++ {
//			b := backing[index]
//			b.style = styles[key]
//			b.value = liner[index]
//			backing[index] = b
//		}
//	}
//	return backing, liner
//}
//
//func getStyle(patterns []TmLanguagePatterns, line []byte) ([][]int, []tcell.Style) {
//	locs := [][]int{}
//	styles := []tcell.Style{}
//	for _, pattern := range patterns {
//		if pattern.Match != "" {
//			reMatch := regexp.MustCompile(pattern.Match)
//			locsAll := reMatch.FindAllIndex(line, -1)
//			if locsAll != nil {
//				for _, loc := range locsAll {
//					styles = append(styles, getThemeColor(pattern.Name))
//					locs = append(locs, loc)
//				}
//
//			}
//		}
//	}
//	return locs, styles
//}
