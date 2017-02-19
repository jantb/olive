package main

import (
	"github.com/jantb/tcell"
	"strings"
)

var theme = Theme{}

type Theme struct {
	Name     string `json:"name"`
	Settings []struct {
		Scope    string `json:"scope"`
		Settings struct {
			FontStyle  string `json:"fontStyle,omitempty"`
			Foreground string `json:"foreground,omitempty"`
			Background string `json:"background,omitempty"`
		} `json:"settings"`
		Name string `json:"name,omitempty"`
	} `json:"settings"`
}

func loadDark() {
	theme = dark
}

var editorStyle tcell.Style

func getEditorStyle() tcell.Style {
	if editorStyle == 0 {
		for _, value := range theme.Settings {
			if value.Scope == "editor" {
				style := tcell.StyleDefault
				switch value.Settings.FontStyle {
				case "bold":
					style = style.Bold(true)
				}
				if value.Settings.Foreground != "" {
					style = style.Foreground(tcell.GetColor(value.Settings.Foreground))
				}
				if value.Settings.Background != "" {
					style = style.Background(tcell.GetColor(value.Settings.Background))
				}
				editorStyle = style
				break
			}
		}
	}

	return editorStyle
}

var syntaxScopes = make(map[string]tcell.Style)

func getThemeColor(scope string) tcell.Style {
	splits := strings.Split(scope, ".")
	for i := len(splits); i > 0; i-- {
		sub := strings.Join(splits[:i], ".")
		s := syntaxScopes[sub]
		if s != 0 {
			return s
		}
	}

	for _, value := range theme.Settings {
		s := ""
		splits := strings.Split(scope, ".")
		for i := len(splits); i > 0; i-- {
			s = strings.Join(splits[:i], ".")
			if value.Scope == s {
				style := getEditorStyle()
				switch value.Settings.FontStyle {
				case "bold":
					style = style.Bold(true)
				}
				if value.Settings.Foreground != "" {
					style = style.Foreground(tcell.GetColor(value.Settings.Foreground))
				}
				if value.Settings.Background != "" {
					style = style.Background(tcell.GetColor(value.Settings.Background))
				}
				syntaxScopes[scope] = style
				return style
			}
		}
	}
	TermMessage("Unknown scope please add to theme:" + scope)
	return tcell.StyleDefault
}
