package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gdamore/tcell"
	"github.com/gdamore/tcell/encoding"
)

var topRow = 0

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
	//backing := [][]Backing{}
	go func() {
		for {
			//drawRuler(s)
			t := time.Now()
			drawBacking(s, drawBuffer(s, buffer))
			s.Show()

			buffer.Insert(1, 0, []byte(fmt.Sprint(time.Now().Sub(t))))
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

func drawBuffer(s tcell.Screen, buffer *Buffer) *[][]Backing {
	w, h := s.Size()
	lines := buffer.GetLines(topRow, topRow+h)
	backing := make([][]Backing, len(lines))
	for i := range backing {
		line := lines[i]

		for _, r := range string(line) {
			b := Backing{}
			b.value = r
			backing[i] = append(backing[i], b)
		}
		for len(backing[i]) < w {
			backing[i] = append(backing[i], Backing{value: ' '})
		}
	}
	return &backing
}

func drawBacking(s tcell.Screen, backing *[][]Backing) {
	if len(previousBacking) != len(*backing) {
		previousBacking = make([][]Backing, len(*backing), len(*backing))
	}

	for ir, row := range *backing {
		if len(previousBacking[ir]) != len(row) {
			previousBacking[ir] = make([]Backing, len(row), len(row))
		}
		x := 0
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

func drawRuler(s tcell.Screen) {
	_, h := s.Size()
	for index := 0; index < h; index++ {
		puts(s, tcell.StyleDefault.
			Foreground(tcell.ColorWhite).
			Background(tcell.ColorDefault), 0, index, fmt.Sprintf("%2d", index+1))
	}
	s.Show()
}
