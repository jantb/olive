package editor

import (
	"github.com/gdamore/tcell"
	"github.com/jantb/olive/rpc"
)

type stylemap map[int]*rpc.DefineStyle

var defaultStyle tcell.Style

var styles = make(stylemap)

type Style struct {
	fg, bg *rpc.RGBA
}

func (sm stylemap) DefineStyle(styledef *rpc.DefineStyle) {
	sm[styledef.ID] = styledef
}

func (sm stylemap) GetTcellStyle(styleId int) tcell.Style {
	styledef := sm[styleId]
	var style tcell.Style
	if styledef.FgColor != 0 {
		r, g, b := styledef.FgColor.ToRGB()
		fg := tcell.NewRGBColor(r, g, b)
		style = style.Foreground(fg)
	}

	if styledef.BgColor != 0 {
		bg := tcell.NewRGBColor(styledef.BgColor.ToRGB())
		style = style.Background(bg)
	}
	return style
}
