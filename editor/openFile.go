package editor

import (
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
	"strconv"
)

type OpenFile struct {
	*tview.Box
	*Editor
	input string
	res   int
}

// NewView returns a new view view primitive.
func (e *Editor) NewOpenFile() *OpenFile {
	return &OpenFile{
		Box:    tview.NewBox().SetBorder(false),
		Editor: e,
	}
}

// Draw draws this primitive onto the screen.
func (g *OpenFile) Draw(screen tcell.Screen) {
	_, bg, _ := defaultStyle.Decompose()
	g.Box.SetBackgroundColor(bg).SetTitle(" Open file ").SetTitleAlign(tview.AlignLeft).SetBorder(true).Draw(screen)
	offx := g.drawText(screen, "Enter file name: ", 0, 0, defaultStyle.Foreground(tcell.ColorLightCyan))
	offx = g.drawText(screen, g.input, offx, 0, defaultStyle.Background(tcell.ColorDarkCyan).Foreground(tcell.ColorYellow))

}

func (g *OpenFile) drawText(screen tcell.Screen, text string, offsetX, offsetY int, style tcell.Style) (offset int) {
	for x, r := range text {
		offset = g.draw(screen, x+offsetX, offsetY, r, style)
	}
	return offset
}

func (g *OpenFile) draw(screen tcell.Screen, x, y int, r rune, style tcell.Style) int {
	xr, yr, _, _ := g.Box.GetInnerRect()
	screen.SetContent(xr+x, yr+y, r, nil, style)
	return x + 1
}

func (g *OpenFile) InputHandler() func(event *tcell.EventKey, setFocus func(p tview.Primitive)) {

	return g.WrapInputHandler(func(key *tcell.EventKey, setFocus func(p tview.Primitive)) {
		switch key.Key() {
		case tcell.KeyEsc:
			g.pages.HidePage("openFile")
			g.focusView()
		case tcell.KeyRune:
			r := key.Rune()
			if len(g.input) < 20 {
				g.input = g.input + string(r)
			}
		case tcell.KeyBackspace2:
			if len(g.input) > 0 {
				g.input = g.input[:len(g.input)-1]
			}
		case tcell.KeyEnter:
			i, e := strconv.Atoi(g.input)
			if e != nil {
				g.pages.HidePage("openFile")
				g.focusView()
				return
			}
			g.res = i
			g.pages.HidePage("openFile")
			g.focusView()
			//	g.view.dataView[g.view.curViewID].OpenFile(g.res - 1)
			g.input = ""
			g.res = 0
		}
	})
}
