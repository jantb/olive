package editor

import (
	"github.com/jantb/olive/rpc"
	"log"

	"github.com/gdamore/tcell"
)

type stylemap map[int]tcell.Style

var defaultStyle tcell.Style

var styles = make(stylemap)

func (sm stylemap) defineStyle(styledef *rpc.DefineStyle) {
	var style tcell.Style

	if styledef.FgColor != 0 {
		r, g, b := styledef.FgColor.ToRGB()
		fg := tcell.NewRGBColor(r, g, b)
		log.Printf("fg: %d, %d, %d", r, g, b)
		style = style.Foreground(fg)
	}

	if styledef.BgColor != 0 {
		bg := tcell.NewRGBColor(styledef.BgColor.ToRGB())
		style = style.Background(bg)
	}

	sm[styledef.ID] = style
}
