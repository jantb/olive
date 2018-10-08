package editor

import (
	"github.com/gdamore/tcell"
	"strings"
	"unicode/utf8"

	runewidth "github.com/mattn/go-runewidth"
)

func ByteWidth(str string, tabSize int) int {
	w := runewidth.StringWidth(str)
	lineIdx := 0
	for _, ch := range str {
		switch ch {
		case '\t':
			ts := tabSize - (lineIdx % tabSize)
			w += ts
			lineIdx += ts
		case '\n':
			lineIdx = 0
		default:
			lineIdx++
		}
	}
	return w
}

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

func GetCursorVisualX(x int, line string) int {
	if x > len(line) {
		x = len(line) - 1
	}

	return ByteWidth(string(line[:x]), tabSize)
}

func Count(b []byte) int {
	return utf8.RuneCountInString(string(b))
}
