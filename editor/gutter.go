package editor

import (
	"github.com/gdamore/tcell"
	"github.com/jantb/olive/ds"
	"github.com/rivo/tview"
)

type Gutter struct {
	*tview.Box
	*Editor

	warning []ds.Position
	error   []ds.Position

	cursorX, cursorY int
}

// NewView returns a new view view primitive.
func (e *Editor) NewGutter() *Gutter {
	e.gutter_width = 2
	return &Gutter{
		Box:    tview.NewBox().SetBorder(false),
		Editor: e,
	}
}

// Draw draws this primitive onto the screen.
func (l *Gutter) Draw(screen tcell.Screen) {
	_, bg, _ := defaultStyle.Decompose()

	l.Box.SetBackgroundColor(bg).Draw(screen)
	invalidBefore := 0
	if l.view.dataView[l.curViewID] != nil {
		invalidBefore = l.view.dataView[l.curViewID].LineCache.InvalidBefore()
	}
	for y, _ := invalidBefore+l.view.offy+1, 0; y <= Min(invalidBefore+l.view.offy+l.view.height, l.view.footer.totalLines); y++ {
		//for _, pos := range l.warning {
		//	if true {
		//
		//	}
		//}
		//l.drawText(screen, fmt.Sprintf("%"+strconv.Itoa(l.linenums_width-1)+"s", strconv.Itoa(y)), yy, 0)
		//yy++
	}
}

func (l *Gutter) drawText(screen tcell.Screen, text string, y, offsetX int) {
	for x, r := range text {
		l.draw(screen, x+offsetX, y, r)
	}
}

func (l *Gutter) draw(screen tcell.Screen, x, y int, r rune) {
	xr, yr, _, _ := l.Box.GetInnerRect()
	screen.SetContent(xr+x, yr+y, r, nil, defaultStyle.Foreground(tcell.ColorLightCyan))
}
