package main

// TODO probably a bug in libui: changing the font away from skia leads to a crash

import (
	"fmt"
	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
)

var mainwin *ui.Window
var (
	attrstr *ui.AttributedString
)

func makeBasicControlsPage() ui.Control {

	entryForm := ui.NewForm()
	entryForm.SetPadded(false)
	entryForm.Append("", ui.NewMultilineEntry(), true)

	return entryForm
}

type areaHandler struct{}

func (areaHandler) Draw(a *ui.Area, p *ui.AreaDrawParams) {
	tl := ui.DrawNewTextLayout(&ui.DrawTextLayoutParams{
		String: attrstr,
		DefaultFont: &ui.FontDescriptor{
			Family:  ui.TextFamily("Fira code"),
			Italic:  ui.TextItalicNormal,
			Size:    ui.TextSize(12),
			Weight:  ui.TextWeightNormal,
			Stretch: ui.TextStretchNormal,
		},
		Width: p.AreaWidth,
		Align: ui.DrawTextAlignLeft,
	})
	defer tl.Free()
	p.Context.Text(tl, 0, 0)
}
func (areaHandler) MouseEvent(a *ui.Area, me *ui.AreaMouseEvent) {
	// do nothing
}

func (areaHandler) MouseCrossed(a *ui.Area, left bool) {
	// do nothing
}

func (areaHandler) DragBroken(a *ui.Area) {
	// do nothing
}

func (areaHandler) KeyEvent(a *ui.Area, ke *ui.AreaKeyEvent) (handled bool) {
	// reject all keys
	return false
}

func setupUI() {

	mainwin = ui.NewWindow("Olive", 640, 480, true)
	mainwin.OnClosing(func(*ui.Window) bool {
		ui.Quit()
		return true
	})
	ui.OnShouldQuit(func() bool {
		mainwin.Destroy()
		return true
	})
	lines := buffer.GetLines(0, 0, 1000, 1000)
	fmt.Println(buffer.Len())
	attrstr = ui.NewAttributedString("")
	for _, value := range lines {
		attrstr.AppendUnattributed(string(value) + "\n")
	}

	area := ui.NewArea(areaHandler{})

	mainwin.SetChild(area)
	mainwin.SetMargined(false)

	mainwin.Show()
}
