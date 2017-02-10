package main

import (
	"fmt"
	"os"

	"strconv"

	"time"

	"github.com/jantb/tcell"
	"github.com/jantb/tcell/encoding"
)

var topRow = 0
var height = 0
var width = 0

//Display renders the editor
func Display(buffer *Buffer) {
	s, e := tcell.NewScreen()
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
	quit := make(chan struct{})
	s.Show()
	w, h := s.Size()
	width = w
	height = h
	backing = make([][]Backing, h-1, h-1)
	for i := range backing {
		backing[i] = make([]Backing, w, w)
	}
	evName := ""
	//backing := [][]Backing{}
	go func() {
		for {
			var offset = drawRuler(topRow, height-1, s)
			var w = width - offset
			lines := buffer.GetLines(topRow, height-1, w)

			t := time.Now()
			drawBuffer(w, topRow, height-1, buffer, s, offset, lines)
			time := time.Now().Sub(t).String()

			puts(s, tcell.StyleDefault.
				Foreground(tcell.ColorWhite).
				Background(tcell.ColorDefault), 0, height-1, "                                                                                                                                                                                                                                                               ")
			puts(s, tcell.StyleDefault.
				Foreground(tcell.ColorWhite).
				Background(tcell.ColorDefault), 0, height-1, time)
			var of = width - len(strconv.Itoa(GetCursor().loc.row)+"/"+strconv.Itoa(buffer.Len()))
			puts(s, tcell.StyleDefault.
				Foreground(tcell.ColorWhite).
				Background(tcell.ColorDefault), of, height-1, strconv.Itoa(GetCursor().loc.row)+"/"+strconv.Itoa(buffer.Len()))
			puts(s, tcell.StyleDefault.
				Foreground(tcell.ColorWhite).
				Background(tcell.ColorDefault), of-2-len(evName), height-1, evName)
			s.Show()
			ev := s.PollEvent()
			switch ev := ev.(type) {
			case *tcell.EventKey:
				switch ev.Key() {
				case tcell.KeyDown:
					//topRow = Min(topRow+1, buffer.Len())
					c := GetCursor()
					c.MoveDown()
				case tcell.KeyUp:
					//topRow = Max(0, topRow-1)
					c := GetCursor()
					c.MoveUp()
				case tcell.KeyLeft:
					c := GetCursor()
					c.MoveLeft()
				case tcell.KeyRight:
					c := GetCursor()
					c.MoveRight()
				case tcell.KeyEscape:
					close(quit)
					return
				case tcell.KeyBackspace2:
					c := GetCursor()
					buffer.Delete(c.loc.row, c.loc.column-1)
					c.MoveLeft()
				case tcell.KeyCtrlS:
					buffer.Save()
				case tcell.KeyRune:
					evName = ev.Name()
					c := GetCursor()
					//Move cursor further if two bytes?
					buffer.Insert(c.loc.row, c.loc.column, []rune(string(ev.Rune())))
					c.MoveRight()
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
var style = tcell.StyleDefault.
	Foreground(tcell.ColorWhite).
	Background(tcell.GetColor("#464742"))

func drawBuffer(w, topRow, h int, buffer *Buffer, s tcell.Screen, offset int, lines [][]rune) {
	c := GetCursor()
	for i := range backing {
		if len(lines)-1 < i {
			continue
		}
		jj := 0
		for _, r := range lines[i] {
			backing[i][jj].value = r
			backing[i][jj].style = style
			jj++
		}
		for index := len(lines[i]); index < w; index++ {
			backing[i][index].value = ' '
			backing[i][index].style = style
		}
	}

	if len(backing) > c.loc.row-topRow && len(backing[c.loc.row-topRow]) > c.loc.column {
		backing[c.loc.row-topRow][c.loc.column].style = tcell.StyleDefault.Reverse(true)
	}

	re := []rune{}
	for ir, row := range backing {
		x := offset
		for _, column := range row {
			if column.value == '\t' {
				s.SetContent(x, ir, ' ', re, style)
				s.SetContent(x+1, ir, ' ', re, style)
				s.SetContent(x+2, ir, ' ', re, style)
				s.SetContent(x+3, ir, ' ', re, style)
				x += 4
				continue
			}
			s.SetContent(x, ir, column.value, nil, column.style)
			x++

			//put(s, column.style, x, ir, column.value)
			//	fmt.Printf("%s", string(column.value))
		}
	}

}

func drawRuler(topRow, h int, s tcell.Screen) int {
	length := strconv.Itoa(len(fmt.Sprintf("%d", topRow+h)))
	ret := len(fmt.Sprintf("%d", topRow+h))
	for index := 0; index < h; index++ {
		puts(s, tcell.StyleDefault.
			Foreground(tcell.ColorDarkSlateGrey).
			Background(tcell.ColorDefault), 0, index, fmt.Sprintf("%"+length+"d ", topRow+index+1))
	}
	return ret + 1
}
