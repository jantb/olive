package main

import (
	"github.com/jantb/tcell"
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

func getThemeColor(scope string) tcell.Style {
	for _, value := range theme.Settings {
		if value.Scope == scope {
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
			return style
		}
	}
	TermMessage("Unknown scope please add to theme:" + scope)
	return tcell.StyleDefault
}
