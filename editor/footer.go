package editor

import (
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
	"strconv"
)

type Footer struct {
	*tview.Box
	*Editor
	totalLines       int
	cursorX, cursorY int
}

// NewView returns a new main view primitive.
func (e *Editor) NewFooter() *Footer {
	return &Footer{
		Box:    tview.NewBox().SetBorder(false),
		Editor: e,
	}
}

// Draw draws this primitive onto the screen.
func (m *Footer) Draw(screen tcell.Screen) {
	_, bg, _ := defaultStyle.Decompose()
	m.Box.SetBackgroundColor(bg).Draw(screen)
	m.drawText(screen, strconv.Itoa(m.totalLines)+
		"L"+" "+strconv.Itoa(m.cursorY)+":"+strconv.Itoa(m.cursorX), 0)
}

func (m *Footer) drawText(screen tcell.Screen, text string, offsetX int) {
	for x, r := range text {
		m.draw(screen, x+offsetX, r)
	}
}

func (m *Footer) draw(screen tcell.Screen, x int, r rune) {
	xr, yr, _, _ := m.Box.GetInnerRect()
	screen.SetContent(xr+x, yr, r, nil, tcell.StyleDefault)
}
