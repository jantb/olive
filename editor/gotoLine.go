package editor

import (
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
	"strconv"
)

type GotoLine struct {
	*tview.Box
	*Editor
	input string
	res   int
}

// NewView returns a new view view primitive.
func (e *Editor) NewGotoLine() *GotoLine {
	return &GotoLine{
		Box:    tview.NewBox().SetBorder(false),
		Editor: e,
	}
}

// Draw draws this primitive onto the screen.
func (g *GotoLine) Draw(screen tcell.Screen) {
	_, bg, _ := defaultStyle.Decompose()
	g.Box.SetBackgroundColor(bg).SetTitle(" Go to line ").SetTitleAlign(tview.AlignLeft).SetBorder(true).Draw(screen)
	offx := g.drawText(screen, "Line: ", 0, defaultStyle.Foreground(tcell.ColorLightCyan))
	offx = g.drawText(screen, g.input, offx, defaultStyle.Background(tcell.ColorDarkCyan).Foreground(tcell.ColorYellow))
}

func (g *GotoLine) drawText(screen tcell.Screen, text string, offsetX int, style tcell.Style) (offset int) {
	for x, r := range text {
		offset = g.draw(screen, x+offsetX, r, style)
	}
	return offset
}

func (g *GotoLine) draw(screen tcell.Screen, x int, r rune, style tcell.Style) int {
	xr, yr, _, _ := g.Box.GetInnerRect()
	screen.SetContent(xr+x, yr, r, nil, style)
	return x + 1
}

func (g *GotoLine) InputHandler() func(event *tcell.EventKey, setFocus func(p tview.Primitive)) {

	return g.WrapInputHandler(func(key *tcell.EventKey, setFocus func(p tview.Primitive)) {
		switch key.Key() {
		case tcell.KeyEsc:
			g.pages.HidePage("gotoLine")
			g.focusView()
		case tcell.KeyRune:
			r := key.Rune()
			_, e := strconv.Atoi(string(r))
			if e != nil {
				break
			}
			if len(g.input) < 12 {
				g.input = g.input + string(r)
			}
		case tcell.KeyBackspace2:
			if len(g.input) > 0 {
				g.input = g.input[:len(g.input)-1]
			}
		case tcell.KeyEnter:
			i, e := strconv.Atoi(g.input)
			if e != nil {
				g.pages.HidePage("gotoLine")
				g.focusView()
				return
			}
			g.res = i
			g.pages.HidePage("gotoLine")
			g.focusView()
			g.view.dataView[g.view.curViewID].GoToLine(g.res - 1)
			g.input = ""
			g.res = 0
		}
	})
}
