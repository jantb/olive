package main

import (
	"strings"

	runewidth "github.com/mattn/go-runewidth"
	"github.com/zyedidia/tcell"
)

func puts(s tcell.Screen, style tcell.Style, x, y int, str string) (length int) {
	i := 0
	var deferred []rune
	dwidth := 0
	str = strings.Replace(str, "\t", "    ", -1)
	for _, r := range str {
		switch runewidth.RuneWidth(r) {
		case 0:
			if len(deferred) == 0 {
				deferred = append(deferred, ' ')
				dwidth = 1
			}
		case 1:
			if len(deferred) != 0 {
				s.SetContent(x+i, y, deferred[0], deferred[1:], style)
				i += dwidth
			}
			deferred = nil
			dwidth = 1
		case 2:
			if len(deferred) != 0 {
				s.SetContent(x+i, y, deferred[0], deferred[1:], style)
				i += dwidth
			}
			deferred = nil
			dwidth = 2
		}
		deferred = append(deferred, r)
	}
	if len(deferred) != 0 {
		s.SetContent(x+i, y, deferred[0], deferred[1:], style)
		i += dwidth
	}
	return i
}

func put(s tcell.Screen, style tcell.Style, x, y int, r rune) (length int) {
	i := 0
	var deferred []rune
	dwidth := 0
	switch runewidth.RuneWidth(r) {
	case 0:
		if len(deferred) == 0 {
			deferred = append(deferred, ' ')
			dwidth = 1
		}
	case 1:
		if len(deferred) != 0 {
			s.SetContent(x+i, y, deferred[0], deferred[1:], style)
			i += dwidth
		}
		deferred = nil
		dwidth = 1
	case 2:
		if len(deferred) != 0 {
			s.SetContent(x+i, y, deferred[0], deferred[1:], style)
			i += dwidth
		}
		deferred = nil
		dwidth = 2
	}
	deferred = append(deferred, r)
	if len(deferred) != 0 {
		s.SetContent(x+i, y, deferred[0], deferred[1:], style)
		i += dwidth
	}
	return i
}
