package editor

import (
	"github.com/gdamore/tcell"
)

type Painter interface {
	SetContent(x int, y int, ch rune, comb []rune, style tcell.Style)
	ShowCursor(x int, y int)
	Size() (int, int)
}

type Viewport struct {
	offx, offy    int
	width, height int
	view          Painter
	viewy, viewx  int
}

func (v *Viewport) GetViewport() (lineStart, lineEnd int) {
	return v.viewy, v.viewy + v.height
}

func (v *Viewport) SetContent(x int, y int, ch rune, comb []rune, style tcell.Style) {
	v.view.SetContent(v.offx+x-v.viewx, v.offy+y-v.viewy, ch, comb, style)
}

func (v *Viewport) MakeVisibleY(y int) {
	if y >= v.viewy+v.height {
		v.viewy = y - (v.height - 1)
	}
	if y >= 0 && y < v.viewy {
		v.viewy = y
	}
	v.ValidateViewY()
}

func (v *Viewport) MakeVisibleX(x int) {
	if x >= v.viewx+v.width {
		v.viewx = x - (v.width - 1)
	}
	if x >= 0 && x < v.viewx {
		v.viewx = x
	}
	v.ValidateViewX()
}

func (v *Viewport) ShowCursor(x int, y int) {
	v.view.ShowCursor(v.offx+x-v.viewx, v.offy+y-v.viewy)
}

func (v *Viewport) ValidateViewY() {
	if v.viewy < 0 {
		v.viewy = 0
	}
}

func (v *Viewport) ValidateViewX() {
	if v.viewx < 0 {
		v.viewx = 0
	}
}

func (v *Viewport) ScrollUp(rows int) {
	v.viewy -= rows
	v.ValidateViewY()
}

func (v *Viewport) ScrollDown(rows int) {
	v.viewy += rows
	v.ValidateViewY()
}

func (v *Viewport) Size() (int, int) {
	return v.width, v.height
}

func (v *Viewport) FillParent() {
	width, height := v.view.Size()
	v.width = width
	v.height = height
}

func (v *Viewport) SetOffsetX(x int) {
	v.offx = x
}

func (v *Viewport) SetOffsetY(y int) {
	v.offy = y
}

func (v *Viewport) SetWidth(w int) {
	v.width = w
}

func (v *Viewport) SetHeight(h int) {
	v.height = h
}

func NewViewport(parent Painter, offx, offy int) *Viewport {
	width, height := parent.Size()
	return &Viewport{
		view:   parent,
		offx:   offx,
		offy:   offy,
		width:  width - offx,
		height: height - offy,
	}
}
