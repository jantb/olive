package main

// TODO probably a bug in libui: changing the font away from skia leads to a crash

import (
	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
)

var mainwin *ui.Window

func makeBasicControlsPage() ui.Control {

	entryForm := ui.NewForm()
	entryForm.SetPadded(false)
	entryForm.Append("", ui.NewMultilineEntry(), true)

	return entryForm
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

	mainwin.SetChild(makeBasicControlsPage())
	mainwin.SetMargined(false)

	mainwin.Show()
}
