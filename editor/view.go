package editor

import (
	"log"

	"github.com/atotto/clipboard"
	"github.com/gdamore/tcell"
	"github.com/jantb/olive/go_plugin"
	"github.com/jantb/olive/rpc"
	"github.com/jantb/olive/xi"
	"github.com/rivo/tview"
)

// View implements the view editor view.
type View struct {
	dataView map[string]*Dataview
	Lines    [][]Block
	*tview.Box
	*Editor
	offy, offx, height, width int
	findstatus                *rpc.FindStatus
}

type Block struct {
	Rune  rune
	Style tcell.Style
}

// NewView returns a new view view primitive.
func NewView() *View {
	view := View{
		Box:   tview.NewBox().SetBorder(false),
		Lines: [][]Block{},
	}
	return &view
}

// Draw draws this primitive onto the screen.
func (v *View) Draw(screen tcell.Screen) {
	_, bg, _ := defaultStyle.Decompose()
	v.Box.SetBackgroundColor(bg).Draw(screen)
	_, _, width, height := v.Box.GetInnerRect()
	v.height = height
	v.width = width
	v.Editor.mutex.Lock()
	defer v.Editor.mutex.Unlock()
	dataview := v.dataView[v.curViewID]
	if dataview == nil {
		return
	}

	lines := dataview.Lines()
	blocksy := [][]Block{}
	offy := v.offy
	offx := v.offx
	blocksy = getBlocks(lines, offy, height, blocksy, offx, width, v)

	v.Lines = blocksy
	v.drawBlocks(screen)
	v.drawCursors(lines, height, screen)
}

func getBlocks(lines []*xi.Line, offy int, height int, blocksy [][]Block, offx int, width int, m *View) [][]Block {
	if len(lines) < offy {
		//return
	}
	for y, line := range lines[offy : offy+height] {
		if line == nil {
			continue
		}
		var blocks []Block
		blocksy = append(blocksy, blocks)

		for x, r := range line.Text[Max(0, Min(offx, len(line.Text)-1)):Max(0, Min(offx+width, len(line.Text)-1))] {
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
			blocksy[y] = append(blocksy[y], Block{Rune: r, Style: style})
		}
	}
	return blocksy
}

func (v *View) drawBlocks(screen tcell.Screen) {
	for y, line := range v.Lines {
		offX := 0
		for x, block := range line {
			if block.Rune == '\t' {
				v.draw(screen, x, y, block)
				v.draw(screen, x+1, y, block)
				v.draw(screen, x+2, y, block)
				v.draw(screen, x+3, y, block)
				offX += 3
				continue
			}
			v.draw(screen, x+offX, y, block)
		}
	}
}

func (v *View) drawCursors(lines []*xi.Line, h int, screen tcell.Screen) {
	for y, line := range lines[v.offy : v.offy+h] {
		if line == nil {
			continue
		}
		for _, cursor := range line.Cursors {
			x := GetCursorVisualX(cursor, line.Text)
			content := v.getContent(screen, x, y)
			content.Style = content.Style.Reverse(true)
			v.draw(screen, x, y, content)
		}
	}
}

func (v *View) draw(screen tcell.Screen, x int, y int, b Block) {

	xMin, yMin, width, height := v.Box.GetInnerRect()
	x = xMin + x - v.offx
	y = yMin + y

	if x < xMin || y < yMin || x >= width+xMin || y >= height+yMin {
		return
	}
	screen.SetContent(x, y, b.Rune, nil, b.Style)
}

func (v *View) getContent(screen tcell.Screen, x int, y int) Block {

	xMin, yMin, width, height := v.Box.GetInnerRect()
	x = xMin + x - v.offx
	y = yMin + y

	if x < xMin || y < yMin || x >= width+xMin || y >= height+yMin {
		return Block{}
	}
	mainc, _, style, _ := screen.GetContent(x, y)
	return Block{Rune: mainc, Style: style}
}

func (v *View) MakeVisible(x, y int) {
	lines := v.dataView[v.curViewID].Lines()
	y -= v.dataView[v.curViewID].LineCache.InvalidBefore()
	x = GetCursorVisualX(x, lines[y].Text)
	_, _, width, height := v.Box.GetInnerRect()

	if y >= v.offy+height {
		v.offy = y - (height - 1)
	}

	if y >= 0 && y < v.offy {
		v.offy = y
	}

	if x >= v.offx+width {
		v.offx = x - (width - 1)
	}
	if x >= 0 && x < v.offx {
		v.offx = x
	}
}

// InputHandler returns the handler for this primitive.
func (v *View) InputHandler() func(event *tcell.EventKey, setFocus func(p tview.Primitive)) {
	return v.WrapInputHandler(func(event *tcell.EventKey, setFocus func(p tview.Primitive)) {
		dataview := v.dataView[v.curViewID]
		ctrl := event.Modifiers()&tcell.ModCtrl != 0
		alt := event.Modifiers()&tcell.ModAlt != 0
		shift := event.Modifiers()&tcell.ModShift != 0
		if !ctrl && !alt && !shift {
			switch event.Key() {
			case tcell.KeyEsc:
				v.findstatus = nil
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
			case tcell.KeyHome:
				dataview.MoveToBeginningOfLine()
			case tcell.KeyEnd:
				dataview.MoveToEndOfLine()
			case tcell.KeyTab:
				dataview.Tab()
			case tcell.KeyBS:
				dataview.DeleteForward()
			case tcell.KeyDelete:
				dataview.DeleteForward()
			case tcell.KeyBackspace2:
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
			case tcell.KeyHome:
				dataview.MoveToBeginningOfLineAndModifySelection()
			case tcell.KeyEnd:
				dataview.MoveToEndOfLineAndModifySelection()
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
				if v.footer.language == "Go" {
					dataview.CancelOperation()
					dataview.SelectAll()
					src := dataview.Copy()
					dataview.Insert(go_plugin.Format(src))
					dataview.CancelOperation()
				} else if v.footer.language == "XML" {
					dataview.CancelOperation()
					dataview.SelectAll()
					src := dataview.Copy()
					dataview.Insert(go_plugin.FormatXml(src))
					dataview.CancelOperation()
				} else if v.footer.language == "JSON" {
					dataview.CancelOperation()
					dataview.SelectAll()
					src := dataview.Copy()
					dataview.Insert(go_plugin.FormatJson(src))
					dataview.CancelOperation()
				}

			}
		}
		if ctrl && !alt && !shift {
			switch event.Key() {
			case tcell.KeyBS:
				dataview.DeleteWordForward()
			case tcell.KeyHome:
				dataview.MoveToBeginningOfDocument()
			case tcell.KeyEnd:
				dataview.MoveToEndOfDocument()
			case tcell.KeyDelete:
				dataview.DeleteWordForward()
			case tcell.KeyBackspace2:
				dataview.DeleteBackward()
			case tcell.KeyLeft:
				dataview.MoveWordLeft()
			case tcell.KeyRight:
				dataview.MoveWordRight()
			case tcell.KeyCtrlS:
				dataview.Save()
			case tcell.KeyCtrlC:
				data := dataview.Copy()
				if data == "" {
					dataview.MoveToBeginningOfLine()
					dataview.MoveToEndOfLineAndModifySelection()
					data = dataview.Copy()
				}
				clipboard.WriteAll(data)
			case tcell.KeyCtrlX:
				clipboard.WriteAll(dataview.Cut())
			case tcell.KeyCtrlA:
				dataview.SelectAll()
			case tcell.KeyCtrlZ:
				dataview.Undo()
			case tcell.KeyCtrlV:
				s, e := clipboard.ReadAll()
				if e != nil {
					log.Println(e)
					return
				}
				dataview.Insert(s)
			case tcell.KeyCtrlQ:
				dataview.Close()
				v.header.path = ""
				v.curViewID = ""
				v.footer.totalLines = 0
				v.footer.cursorX = 0
				v.footer.cursorY = 0
				v.focusFileselector()
			case tcell.KeyCtrlD:
				dataview.DuplicateLine()
			default:
				log.Println(event.Name())
			}
		}
		if !ctrl && alt {
			switch event.Name() {
			case "Alt+Up":
				dataview.AddSelectionAbove()
			case "Alt+Rune[j]":
				if v.findstatus == nil {
					s := dataview.Copy()
					if s == "" {
						dataview.MoveWordLeft()
						dataview.MoveWordRightAndModifySelection()
						s = dataview.Copy()
					}
					dataview.Find(s, false, false, false)
				}
				dataview.FindNext(true, true, "add")
			case "Alt+Down":
				dataview.AddSelectionBelow()
			case "Alt+Rune[0]":
				v.focusFileselector()
			default:
				log.Println(event.Name())
			}
		}
		log.Println(event.Name())
	})
}
