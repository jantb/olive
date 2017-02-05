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
	previousBacking = make([][]Backing, h-1, h-1)
	for i := range previousBacking {
		previousBacking[i] = make([]Backing, w, w)
	}
	//backing := [][]Backing{}
	go func() {
		for {
			var offset = drawRuler(topRow, h, s)
			var w = width - offset
			t := time.Now()
			drawBuffer(w, topRow, height-1, buffer, s, offset)
			time := time.Now().Sub(t).String()
			puts(s, tcell.StyleDefault.
				Foreground(tcell.ColorWhite).
				Background(tcell.ColorDefault), 0, height-1, "                                                                                            ")
			puts(s, tcell.StyleDefault.
				Foreground(tcell.ColorWhite).
				Background(tcell.ColorDefault), 0, height-1, time)
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
					buffer.RemoveRow(0)
					buffer.Insert(0, 0, []byte(ev.Name()))
				}
			case *tcell.EventResize:
				s.Sync()
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

var previousBacking = [][]Backing{}
var backing = [][]Backing{}

func drawBuffer(w, topRow, h int, buffer *Buffer, s tcell.Screen, offset int) {
	lines := buffer.GetLines(topRow, h, w)

	for i := range backing {
		for j, r := range string(lines[i]) {
			backing[i][j].value = r
		}
		for index := len(lines[i]); index < w; index++ {
			backing[i][index].value = ' '
		}
	}

	for ir, row := range backing {
		x := offset
		for _, column := range row {
			// buffer equal does not matter, it's whats painted that does
			//if column == previousBacking[ir][x] {
			//continue
			//	}
			//previousBacking[ir][ic] = column
			x += puts(s, tcell.StyleDefault.
				Foreground(tcell.ColorWhite).
				Background(tcell.ColorDefault), x, ir, string(column.value))
		}
	}

}

func drawBacking(offset, w, h int, s tcell.Screen, backing *[][]Backing) {
	if len(previousBacking) != len(*backing) {
		previousBacking = make([][]Backing, len(*backing))
	}

	for ir, row := range *backing {
		if len(previousBacking[ir]) != len(row) {
			previousBacking[ir] = make([]Backing, len(row), len(row))
		}
		x := offset
		for ic, column := range row {
			if column == previousBacking[ir][ic] {
				continue
			}
			previousBacking[ir][ic] = column
			x += puts(s, tcell.StyleDefault.
				Foreground(tcell.ColorWhite).
				Background(tcell.ColorDefault), x, ir, string(column.value))
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
