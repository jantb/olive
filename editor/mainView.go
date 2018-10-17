package editor

import (
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
	"log"
)

// MainView implements the main editor view.
type MainView struct {
	dataView map[string]*Dataview
	Lines    [][]Block
	*tview.Box
	*Edit
}

type Block struct {
	Rune  rune
	Style tcell.Style
}

// NewMainView returns a new main view primitive.
func NewMainView() *MainView {
	return &MainView{
		Box:   tview.NewBox().SetBorder(true),
		Lines: [][]Block{},
	}
}

// Draw draws this primitive onto the screen.
func (m *MainView) Draw(screen tcell.Screen) {
	_, bg, _ := defaultStyle.Decompose()
	m.Box.SetBackgroundColor(bg).Draw(screen)
	dataview := m.dataView[m.curViewID]
	if dataview == nil {
		return
	}
	lines := dataview.Lines()
	m.Lines = [][]Block{}
	for y, line := range lines {
		var blocks []Block
		m.Lines = append(m.Lines, blocks)
		for x, r := range line.Text {
			var style = defaultStyle
			if line.StyleIds[x] >= 2 {
				fg, _, _ := styles[line.StyleIds[x]].Decompose()
				style = style.Foreground(fg)
			}
			m.Lines[y] = append(m.Lines[y], Block{Rune: r, Style: style})
		}
	}

	for y, line := range m.Lines {
		for x, block := range line {
			m.draw(screen, x, y, block)
		}
	}

	// Draw cursors
	for y, line := range lines {
		for _, cursor := range line.Cursors {
			content := m.getContent(screen, cursor, y)
			content.Style = content.Style.Reverse(true)
			m.draw(screen, cursor, y, content)
		}
	}
}

func (m *MainView) draw(screen tcell.Screen, x int, y int, b Block) {

	xMin, yMin, width, height := m.Box.GetInnerRect()
	x = xMin + x
	y = yMin + y

	if x < xMin || y < yMin || x >= width+xMin || y >= height+yMin {
		return
	}
	screen.SetContent(x, y, b.Rune, nil, b.Style)
}

func (m *MainView) getContent(screen tcell.Screen, x int, y int) Block {

	xMin, yMin, width, height := m.Box.GetInnerRect()
	x = xMin + x
	y = yMin + y

	if x < xMin || y < yMin || x >= width+xMin || y >= height+yMin {
		return Block{}
	}
	mainc, _, style, _ := screen.GetContent(x, y)
	return Block{Rune: mainc, Style: style}
}

// InputHandler returns the handler for this primitive.
func (m *MainView) InputHandler() func(event *tcell.EventKey, setFocus func(p tview.Primitive)) {
	return m.WrapInputHandler(func(event *tcell.EventKey, setFocus func(p tview.Primitive)) {
		dataview := m.dataView[m.curViewID]
		ctrl := event.Modifiers()&tcell.ModCtrl != 0
		alt := event.Modifiers()&tcell.ModAlt != 0
		if !ctrl && !alt {
			switch event.Key() {
			case tcell.KeyUp:
				dataview.MoveUp()
			case tcell.KeyDown:
				dataview.MoveDown()
			case tcell.KeyLeft:
				dataview.MoveLeft()
			case tcell.KeyEnter:
				dataview.Newline()
			case tcell.KeyRight:
				dataview.MoveRight()
			case tcell.KeyRune:
				dataview.Insert(string(event.Rune()))
			case tcell.KeyTab:
				dataview.Tab()
			case tcell.KeyBS:
				dataview.DeleteForward()
			case tcell.KeyDEL:
				dataview.DeleteBackward()
			case tcell.KeyPgUp:
				dataview.ScrollPageUp()
			case tcell.KeyPgDn:
				dataview.ScrollPageDown()
			default:
				log.Println(event.Name())
			}
		}
		if ctrl && !alt {
			switch event.Key() {
			case tcell.KeyCtrlS:
				dataview.Save()
			case tcell.KeyCtrlQ:
				dataview.Save()
				dataview.Close()
				m.curViewID = ""
				m.focusFileselector()
			default:
				log.Println(event.Name())
			}
		}
		if !ctrl && alt {
			switch event.Name() {
			case "Alt+Rune[0]":
				m.focusFileselector()
			default:
				log.Println(event.Name())
			}
		}
		if ctrl && alt {
			switch event.Key() {

			default:
				log.Println(event.Name())
			}
		}
		//dataview.Save()
	})
}
