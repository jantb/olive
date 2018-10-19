package editor

import (
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

type Header struct {
	*tview.Box
	*Editor
	path     string
	pristine bool
}

func (e *Editor) NewHeader() *Header {
	return &Header{
		Box:    tview.NewBox().SetBorder(false),
		Editor: e,
	}
}

// Draw draws this primitive onto the screen.
func (m *Header) Draw(screen tcell.Screen) {
	_, bg, _ := defaultStyle.Decompose()
	m.Box.SetBackgroundColor(bg).Draw(screen)
	p := "*"
	if m.pristine || len(m.path) == 0 {
		p = ""
	}
	m.drawText(screen, m.path+p, 0)
}

func (m *Header) drawText(screen tcell.Screen, text string, offsetX int) {
	for x, r := range text {
		m.draw(screen, x+offsetX, r)
	}
}

func (m *Header) draw(screen tcell.Screen, x int, r rune) {
	xr, yr, _, _ := m.Box.GetInnerRect()
	screen.SetContent(xr+x, yr, r, nil, defaultStyle)
}
