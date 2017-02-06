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
			var offset = drawRuler(topRow, h-1, s)
			var w = width - offset
			lines := buffer.GetLines(topRow, h-1, w)

			t := time.Now()
			drawBuffer(w, topRow, height-1, buffer, s, offset, lines)
			time := time.Now().Sub(t).String()

			puts(s, tcell.StyleDefault.
				Foreground(tcell.ColorWhite).
				Background(tcell.ColorDefault), 0, height-1, "                                                                                                                                                                                                                                                               ")
			puts(s, tcell.StyleDefault.
				Foreground(tcell.ColorWhite).
				Background(tcell.ColorDefault), 0, height-1, time)
			puts(s, tcell.StyleDefault.
				Foreground(tcell.ColorWhite).
				Background(tcell.ColorDefault), width-len(strconv.Itoa(buffer.Len())), height-1, strconv.Itoa(buffer.Len()))
			puts(s, tcell.StyleDefault.
				Foreground(tcell.ColorWhite).
				Background(tcell.ColorDefault), width-len(strconv.Itoa(buffer.Len()))-2-len(evName), height-1, evName)
			s.Show()
			ev := s.PollEvent()
			switch ev := ev.(type) {
			case *tcell.EventKey:
				switch ev.Key() {
				case tcell.KeyDown:
					topRow = Min(topRow+1, buffer.Len())
				case tcell.KeyUp:
					topRow = Max(0, topRow-1)
				case tcell.KeyLeft:
					c := GetCursor()
					c.MoveLeft()
				case tcell.KeyRight:
					c := GetCursor()
					c.MoveRight()
				case tcell.KeyEscape:
					close(quit)
					return
				case tcell.KeyRune:
					evName = ev.Name()
					c := GetCursor()
					buffer.Insert(c.loc.row, c.loc.column, []byte(string(ev.Rune())))
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
	Background(tcell.ColorDefault)

func drawBuffer(w, topRow, h int, buffer *Buffer, s tcell.Screen, offset int, lines [][]byte) {

	for i := range backing {
		if len(lines)-1 < i {
			continue
		}
		for j, r := range string(lines[i]) {
			backing[i][j].value = r
			backing[i][j].style = tcell.StyleDefault
		}
		for index := len(lines[i]); index < w; index++ {
			backing[i][index].value = ' '
			backing[i][index].style = tcell.StyleDefault
		}
	}
	c := GetCursor()

	backing[c.loc.row][c.loc.column].style = tcell.StyleDefault.Reverse(true)
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
			put(s, column.style, x, ir, column.value)
			x++
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
