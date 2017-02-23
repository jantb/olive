package main

import (
	"fmt"
	"os"

	"strconv"

	"time"

	"github.com/golang/groupcache/lru"
	"github.com/jantb/tcell"
	"github.com/jantb/tcell/encoding"
)

var topRow = 0
var leftColumn = 0
var height = 0
var width = 0
var offset = 0
var s tcell.Screen

func createDisplay() {
	loadDark()
	os.Setenv("TERM", "xterm-truecolor")
	sc, e := tcell.NewScreen()
	s = sc
	if e != nil {
		fmt.Fprintf(os.Stderr, "%v\n", e)
		os.Exit(1)
	}
	encoding.Register()

	if e = s.Init(); e != nil {
		fmt.Fprintf(os.Stderr, "%v\n", e)
		os.Exit(1)
	}

	s.SetStyle(tcell.StyleDefault.
		Foreground(tcell.ColorWhite).
		Background(tcell.ColorDefault))
	s.Clear()
	s.Show()
	w, h := s.Size()
	width = w
	height = h
	backing = make([][]Backing, h-1, h-1)
	for i := range backing {
		backing[i] = make([]Backing, w, w)
	}
}

//Display renders the editor
func Display(buffer *Buffer) {
	createDisplay()
	quit := make(chan struct{})
	evName := ""
	go func() {
		for {
			offset = drawRuler(topRow, height-1, s)
			var w = width - offset + 1
			lines := buffer.GetLines(topRow, leftColumn, height-1, w)
			t := time.Now()
			drawBuffer(topRow, s, offset, lines)
			time2 := time.Now().Sub(t).String()

			puts(s, tcell.StyleDefault.
				Foreground(tcell.ColorWhite).
				Background(tcell.ColorDefault), 0, height-1, "                                                                                                                                                                                                                                                               ")
			puts(s, tcell.StyleDefault.
				Foreground(tcell.ColorWhite).
				Background(tcell.ColorDefault), 0, height-1, buffer.filename)

			puts(s, tcell.StyleDefault.
				Foreground(tcell.ColorWhite).
				Background(tcell.ColorDefault), len(buffer.filename)+1, height-1, time2)

			var text = strconv.Itoa(GetCursor().loc.row+1) + ":" + strconv.Itoa(GetCursor().loc.column+1) +
				"/" + strconv.Itoa(buffer.Len()) + ":" + strconv.Itoa(len(buffer.GetLine(GetCursor().loc.row)))

			puts(s, tcell.StyleDefault.
				Foreground(tcell.ColorWhite).
				Background(tcell.ColorDefault), width-len(text), height-1, text)

			puts(s, tcell.StyleDefault.
				Foreground(tcell.ColorWhite).
				Background(tcell.ColorDefault), width-len(text)-2-len(evName), height-1, evName)
			s.Show()
			ev := s.PollEvent()
			switch ev := ev.(type) {
			case *tcell.EventKey:
				switch ev.Key() {
				case tcell.KeyDown:
					c := GetCursor()
					c.MoveDown()
				case tcell.KeyUp:
					c := GetCursor()
					c.MoveUp()
				case tcell.KeyLeft:
					c := GetCursor()
					c.MoveLeft()
				case tcell.KeyRight:
					c := GetCursor()
					c.MoveRight()
				case tcell.KeyPgDn:
					c := GetCursor()
					c.MovePageDown()
				case tcell.KeyPgUp:
					c := GetCursor()
					c.MovePageUp()
				case tcell.KeyEscape:
					close(quit)
					return
				case tcell.KeyCtrlD:
					c := GetCursor()
					buffer.Insert(c.loc.row, 0, buffer.GetLine(c.loc.row))
					c.MoveDown()
				case tcell.KeyCtrlY:
					c := GetCursor()
					buffer.RemoveRow(c.loc.row)
				case tcell.KeyDelete:
					c := GetCursor()
					buffer.Delete(c.loc.row, c.loc.column)
				case tcell.KeyBackspace2:
					c := GetCursor()
					l := buffer.Backspace(c.loc.row, c.loc.column)
					if c.loc.column == 0 && c.loc.row > 0 {
						c.MoveUp()
						c.MoveToEndOfLine()
						for index := 0; index < l; index++ {
							c.MoveLeft()
						}
					} else {
						c.MoveLeft()
					}
				case tcell.KeyCtrlS:
					buffer.Save()
				case tcell.KeyTab:
					c := GetCursor()
					buffer.Insert(c.loc.row, c.loc.column, []rune(string('\t')))
					c.MoveRight()
				case tcell.KeyRune:
					evName = ev.Name()
					c := GetCursor()
					buffer.Insert(c.loc.row, c.loc.column, []rune(string(ev.Rune())))
					c.MoveRight()
				case tcell.KeyEnter:
					c := GetCursor()
					//buffer.Insert(c.loc.row, c.loc.column, []rune(string(ev.Rune())))
					buffer.InsertEnter(c.loc.row, c.loc.column)
					c.MoveDown()
					c.MoveStartOfLine()
				default:
					evName = ev.Name()
				}

			case *tcell.EventResize:
				s.Sync()
				w, h := s.Size()
				width = w
				height = h
				backing = make([][]Backing, h-1, h-1)
				for i := range backing {
					backing[i] = make([]Backing, w, w)
				}
			}
		}
	}()

	<-quit

	s.Fini()
}

// Backing backs the display so we don't overwrite runes unnessecary
type Backing struct {
	value rune
	style tcell.Style
}

var backing = [][]Backing{}
var cache = lru.New(1000)

func drawBuffer(topRow int, s tcell.Screen, offset int, lines [][]rune) {
	var style = getThemeColor("editor")
	c := GetCursor()
	for i := range backing {
		if len(lines)-1 < i {
			continue
		}
		jj := 0
		key := string(lines[i])
		s := []Backing{}
		if val, ok := cache.Get(key); !ok {
			val := syntax(lines[i], buffer.filename)
			cache.Add(key, val)
			s = val
		} else {
			s = val.([]Backing)
		}

		for o, r := range lines[i] {
			backing[i][jj].value = r
			backing[i][jj].style = s[o].style
			jj++
		}
		for index := len(lines[i]); index < width; index++ {
			backing[i][index].value = ' '
			backing[i][index].style = style
		}
	}

	if len(backing) > c.loc.row-topRow && len(backing[c.loc.row-topRow]) > c.loc.column-leftColumn {
		for i := range backing[c.loc.row-topRow] {
			backing[c.loc.row-topRow][i].style = backing[c.loc.row-topRow][i].style.Background(tcell.GetColor("#282828"))
		}
		// No cursor on enter
		if c.loc.column != -1 {
			if backing[c.loc.row-topRow][c.loc.column-leftColumn].value == '\n' {
				backing[c.loc.row-topRow][c.loc.column-leftColumn+1].style = tcell.StyleDefault.Reverse(true)
			} else {
				backing[c.loc.row-topRow][c.loc.column-leftColumn].style = tcell.StyleDefault.Reverse(true)
			}
		}
	}

	re := []rune{}
	for ir, row := range backing {
		x := offset
		for _, column := range row {
			if column.value == '\t' {
				s.SetContent(x, ir, ' ', re, column.style)
				s.SetContent(x+1, ir, ' ', re, column.style)
				s.SetContent(x+2, ir, ' ', re, column.style)
				s.SetContent(x+3, ir, ' ', re, column.style)
				x += 4
				continue
			}
			if column.value == '\n' {
				continue
			}
			s.SetContent(x, ir, column.value, nil, column.style)
			x++
		}
	}
}

func drawRuler(topRow, h int, s tcell.Screen) int {
	length := strconv.Itoa(len(fmt.Sprintf("%d", topRow+h)))
	ret := len(fmt.Sprintf("%d", topRow+h))
	style := getThemeColor("editor.ruler")
	for index := 0; index < h; index++ {
		puts(s, style, 0, index, fmt.Sprintf("%"+length+"d ", topRow+index+1))
	}
	return ret + 1
}
