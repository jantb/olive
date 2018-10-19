package editor

import (
	"fmt"
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
	"strconv"
)

type Linenums struct {
	*tview.Box
	*Editor
	language         string
	cursorX, cursorY int
}

// NewView returns a new view view primitive.
func (e *Editor) NewLinenums() *Linenums {
	return &Linenums{
		Box:    tview.NewBox().SetBorder(false),
		Editor: e,
	}
}

// Draw draws this primitive onto the screen.
func (l *Linenums) Draw(screen tcell.Screen) {
	_, bg, _ := defaultStyle.Decompose()
	l.linenums_width = len(strconv.Itoa(l.view.footer.totalLines)) + 1
	l.updateColumnWidths()
	l.Box.SetBackgroundColor(bg).Draw(screen)
	invalidBefore := 0
	if l.view.dataView[l.curViewID] != nil {
		invalidBefore = l.view.dataView[l.curViewID].LineCache.InvalidBefore()
	}
	for y, yy := invalidBefore+l.view.offy+1, 0; y <= Min(invalidBefore+l.view.offy+l.view.height, l.view.footer.totalLines); y++ {
		l.drawText(screen, fmt.Sprintf("%"+strconv.Itoa(l.linenums_width-1)+"s", strconv.Itoa(y)), yy, 0)
		yy++
	}
}

func (l *Linenums) drawText(screen tcell.Screen, text string, y, offsetX int) {
	for x, r := range text {
		l.draw(screen, x+offsetX, y, r)
	}
}

func (l *Linenums) draw(screen tcell.Screen, x, y int, r rune) {
	xr, yr, _, _ := l.Box.GetInnerRect()
	screen.SetContent(xr+x, yr+y, r, nil, defaultStyle.Foreground(tcell.ColorLightCyan))
}
