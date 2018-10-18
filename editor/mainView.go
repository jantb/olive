package editor

import (
	"github.com/atotto/clipboard"
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
	offy, offx int
}

type Block struct {
	Rune  rune
	Style tcell.Style
}

// NewMainView returns a new main view primitive.
func NewMainView() *MainView {
	view := MainView{
		Box:   tview.NewBox().SetBorder(true),
		Lines: [][]Block{},
	}
	return &view
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
	for y, line := range lines[m.offy:] {
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
	for y, line := range lines[m.offy:] {
		for _, cursor := range line.Cursors {
			x := GetCursorVisualX(cursor, line.Text)
			content := m.getContent(screen, cursor, y)
			content.Style = content.Style.Reverse(true)
			m.draw(screen, x, y, content)
		}
	}
}

func (m *MainView) draw(screen tcell.Screen, x int, y int, b Block) {

	xMin, yMin, width, height := m.Box.GetInnerRect()
	x = xMin + x - m.offx
	y = yMin + y

	if x < xMin || y < yMin || x >= width+xMin || y >= height+yMin {
		return
	}
	screen.SetContent(x, y, b.Rune, nil, b.Style)
}

func (m *MainView) getContent(screen tcell.Screen, x int, y int) Block {

	xMin, yMin, width, height := m.Box.GetInnerRect()
	x = xMin + x - m.offx
	y = yMin + y

	if x < xMin || y < yMin || x >= width+xMin || y >= height+yMin {
		return Block{}
	}
	mainc, _, style, _ := screen.GetContent(x, y)
	return Block{Rune: mainc, Style: style}
}

func (m *MainView) MakeVisible(x, y int) {
	_, _, width, height := m.Box.GetInnerRect()

	if y >= m.offy+height {
		m.offy = y - (height - 1)
	}

	if y >= 0 && y < m.offy {
		m.offy = y
	}

	if x >= m.offx+width {
		m.offx = x - (width - 1)
	}
	if x >= 0 && x < m.offx {
		m.offx = x
	}
}

// InputHandler returns the handler for this primitive.
func (m *MainView) InputHandler() func(event *tcell.EventKey, setFocus func(p tview.Primitive)) {
	return m.WrapInputHandler(func(event *tcell.EventKey, setFocus func(p tview.Primitive)) {
		dataview := m.dataView[m.curViewID]
		ctrl := event.Modifiers()&tcell.ModCtrl != 0
		alt := event.Modifiers()&tcell.ModAlt != 0
		shift := event.Modifiers()&tcell.ModShift != 0
		log.Println(shift)
		if !ctrl && !alt && !shift {
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
		if !ctrl && !alt && shift {
			switch event.Key() {
			case tcell.KeyRune:
				dataview.Insert(string(event.Rune()))
			default:
				log.Println(event.Name())
			}
			switch event.Name() {
			case "Shift+Right":
				dataview.MoveRightAndModifySelection()
			case "Shift+Left":
				dataview.MoveLeftAndModifySelection()
			case "Shift+Up":
				dataview.MoveUpAndModifySelection()
			case "Shift+Down":
				dataview.MoveDownAndModifySelection()
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
			case tcell.KeyCtrlD:
				dataview.DuplicateLine()
			case tcell.KeyCtrlV:
				s, e := clipboard.ReadAll()
				if e != nil {
					log.Println(e)
					return
				}
				dataview.Insert(s)
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
		log.Println(event.Name())
		//dataview.Save()
	})
}
