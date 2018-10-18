package editor

import (
	"github.com/atotto/clipboard"
	"github.com/gdamore/tcell"
	"github.com/jantb/olive/goPlugin"
	"github.com/rivo/tview"
	"log"
)

// View implements the main editor view.
type View struct {
	dataView map[string]*Dataview
	Lines    [][]Block
	*tview.Box
	*Editor
	offy, offx int
}

type Block struct {
	Rune  rune
	Style tcell.Style
}

// NewView returns a new main view primitive.
func NewView() *View {
	view := View{
		Box:   tview.NewBox().SetBorder(false),
		Lines: [][]Block{},
	}
	return &view
}

// Draw draws this primitive onto the screen.
func (m *View) Draw(screen tcell.Screen) {
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
			if line.StyleIds[x] != nil {
				for _, value := range line.StyleIds[x] {
					s := styles[value]
					if value == 0 {
						s = s.Background(tcell.NewRGBColor(m.Editor.theme.Selection.ToRGB()))
					}

					fg, bg, _ := s.Decompose()

					if fg != tcell.ColorDefault {
						style = style.Foreground(fg)
					}
					if bg != tcell.ColorDefault {
						style = style.Background(bg)
					}
				}
			}
			m.Lines[y] = append(m.Lines[y], Block{Rune: r, Style: style})
		}
	}

	for y, line := range m.Lines {
		offX := 0
		for x, block := range line {
			if block.Rune == '\t' {
				m.draw(screen, x, y, block)
				m.draw(screen, x+1, y, block)
				m.draw(screen, x+2, y, block)
				m.draw(screen, x+3, y, block)
				offX += 3
				continue
			}
			m.draw(screen, x+offX, y, block)
		}
	}

	// Draw cursors
	for y, line := range lines[m.offy:] {
		for _, cursor := range line.Cursors {
			x := GetCursorVisualX(cursor, line.Text)
			content := m.getContent(screen, x, y)
			content.Style = content.Style.Reverse(true)
			m.draw(screen, x, y, content)
		}
	}
}

func (m *View) draw(screen tcell.Screen, x int, y int, b Block) {

	xMin, yMin, width, height := m.Box.GetInnerRect()
	x = xMin + x - m.offx
	y = yMin + y

	if x < xMin || y < yMin || x >= width+xMin || y >= height+yMin {
		return
	}
	screen.SetContent(x, y, b.Rune, nil, b.Style)
}

func (m *View) getContent(screen tcell.Screen, x int, y int) Block {

	xMin, yMin, width, height := m.Box.GetInnerRect()
	x = xMin + x - m.offx
	y = yMin + y

	if x < xMin || y < yMin || x >= width+xMin || y >= height+yMin {
		return Block{}
	}
	mainc, _, style, _ := screen.GetContent(x, y)
	return Block{Rune: mainc, Style: style}
}

func (m *View) MakeVisible(x, y int) {
	lines := m.dataView[m.curViewID].Lines()
	if len(lines) <= y {
		return
	}
	x = GetCursorVisualX(x, lines[y].Text)
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
func (m *View) InputHandler() func(event *tcell.EventKey, setFocus func(p tview.Primitive)) {
	return m.WrapInputHandler(func(event *tcell.EventKey, setFocus func(p tview.Primitive)) {
		dataview := m.dataView[m.curViewID]
		ctrl := event.Modifiers()&tcell.ModCtrl != 0
		alt := event.Modifiers()&tcell.ModAlt != 0
		shift := event.Modifiers()&tcell.ModShift != 0
		if !ctrl && !alt && !shift {
			switch event.Key() {
			case tcell.KeyEsc:
				dataview.CancelOperation()
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
			case tcell.KeyPgUp:
				dataview.ScrollPageUpAndModifySelection()
			case tcell.KeyPgDn:
				dataview.ScrollPageDownAndModifySelection()
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
		if ctrl && alt {
			switch event.Name() {
			case "Alt+Ctrl+Z":
				dataview.Redo()
			case "Alt+Ctrl+L":
				if m.footer.language == "Go" {
					dataview.CancelOperation()
					dataview.SelectAll()
					src := dataview.Copy()
					dataview.Insert(goPlugin.Format(src))
					dataview.CancelOperation()
				} else if m.footer.language == "XML" {
					dataview.CancelOperation()
					dataview.SelectAll()
					src := dataview.Copy()
					dataview.Insert(goPlugin.FormatXml(src))
					dataview.CancelOperation()
				} else if m.footer.language == "JSON" {
					dataview.CancelOperation()
					dataview.SelectAll()
					src := dataview.Copy()
					dataview.Insert(goPlugin.FormatJson(src))
					dataview.CancelOperation()
				}

			}
		}
		if ctrl && !alt && !shift {
			switch event.Key() {
			case tcell.KeyCtrlS:
				dataview.Save()
			case tcell.KeyCtrlA:
				dataview.SelectAll()
			case tcell.KeyCtrlZ:
				dataview.Undo()
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
			case "Alt+Up":
				dataview.AddSelectionAbove()
			case "Alt+Down":
				dataview.AddSelectionBelow()
			case "Alt+Rune[0]":
				m.focusFileselector()
			case "Alt+Rune[c]":
				clipboard.WriteAll(dataview.Copy())
			case "Alt+Rune[x]":
				clipboard.WriteAll(dataview.Cut())
			case "Alt+Rune[v]":
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
		if ctrl && alt {
			switch event.Key() {

			default:
				log.Println(event.Name())
			}
		}
		log.Println(event.Name())
	})
}
