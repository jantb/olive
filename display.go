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
			s.Show()
			ev := s.PollEvent()
			switch ev := ev.(type) {
			case *tcell.EventKey:
				switch ev.Key() {
				case tcell.KeyDown:
					topRow++
				case tcell.KeyUp:
					topRow = Max(0, topRow-1)
				case tcell.KeyEscape, tcell.KeyEnter:
					close(quit)
					return
				default:
					puts(s, tcell.StyleDefault.
						Foreground(tcell.ColorWhite).
						Background(tcell.ColorDefault), width-len(strconv.Itoa(buffer.Len()))-2-len(ev.Name()), height-1, ev.Name())
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
		}
		for index := len(lines[i]); index < w; index++ {
			backing[i][index].value = ' '
		}
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
			s.SetContent(x, ir, column.value, re, style)
			x++
		}
	}

}

func drawRuler(topRow, h int, s tcell.Screen) int {
	length := strconv.Itoa(len(fmt.Sprintf("%d", topRow+h)))
	ret := len(fmt.Sprintf("%d", topRow+h))
	for index := 0; index < h; index++ {
		puts(s, tcell.StyleDefault.
			Foreground(tcell.ColorWhite).
			Background(tcell.ColorDefault), 0, index, fmt.Sprintf("%"+length+"d ", topRow+index+1))
	}
	return ret + 1
}
