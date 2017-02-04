package main

import (
	"fmt"
	"os"

	"github.com/gdamore/tcell"
	"github.com/gdamore/tcell/encoding"
)

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
			//	t := time.Now()
			drawBacking(s, drawBuffer(s, buffer))
			s.Show()
			//buffer.r.Insert(0, []byte(fmt.Sprint(time.Now().Sub(t))))
			ev := s.PollEvent()
			switch ev := ev.(type) {
			case *tcell.EventKey:
				switch ev.Key() {
				case tcell.KeyEscape, tcell.KeyEnter:
					close(quit)
					return
				default:
					puts(s, tcell.StyleDefault.
						Foreground(tcell.ColorWhite).
						Background(tcell.ColorDefault), 10, 0, "               ")

					puts(s, tcell.StyleDefault.
						Foreground(tcell.ColorWhite).
						Background(tcell.ColorDefault), 10, 0, ev.Name())

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

	lines := buffer.getLines(0, 10)
	backing := make([][]Backing, len(lines))
	_, height := s.Size()
	for i := range backing {
		if i > height {
			break
		}
		line := lines[i]
		for _, r := range string(line) {
			b := Backing{}
			b.value = r
			backing[i] = append(backing[i], b)
		}
	}
	return &backing
}

func drawBacking(s tcell.Screen, backing *[][]Backing) {
	if len(previousBacking) != len(*backing) {
		previousBacking = make([][]Backing, len(*backing), len(*backing))
	}

	for ir, row := range *backing {
		prevRow := previousBacking[ir]
		if len(prevRow) != len(row) {
			prevRow = make([]Backing, len(row), len(row))
		}
		x := 0
		for ic, column := range row {
			if column == prevRow[ic] {
				continue
			}
			prevRow[ic] = column
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
