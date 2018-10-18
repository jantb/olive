package editor

import (
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

type FooterTree struct {
	*tview.Box
	*Edit
}

// NewMainView returns a new main view primitive.
func (e *Edit) NewFooterTree() *FooterTree {
	return &FooterTree{
		Box:  tview.NewBox().SetBorder(false),
		Edit: e,
	}
}

// Draw draws this primitive onto the screen.
func (m *FooterTree) Draw(screen tcell.Screen) {
	_, bg, _ := defaultStyle.Decompose()
	m.Box.SetBackgroundColor(bg).Draw(screen)

}

func (m *FooterTree) draw(screen tcell.Screen, x int, r rune) {
	xr, yr, _, _ := m.Box.GetInnerRect()
	screen.SetContent(xr+x, yr, r, nil, tcell.StyleDefault)
}
