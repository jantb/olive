package editor

import (
	"github.com/atotto/clipboard"
	"github.com/jantb/olive/rpc"
	"log"
	"strconv"

	"github.com/gdamore/tcell"
)

const tabSize = 4

// TODO: Move this to RPC
type requestLines struct {
	Method string `json:"method"`
	Params []int  `json:"params"`
	ViewID string `json:"view_id"`
}

type View struct {
	*LineCache
	*InputHandler

	ID        string
	view      *Viewport
	gutter    *Viewport
	xi        *rpc.Connection
	ViewID    string
	lineStart int
	lineEnd   int
}

func ralign(str string, width int) string {
	diff := width - len(str)
	res := ""
	for i := 0; i < diff; i++ {
		res += " "
	}

	res += str
	return res
}

func NewView(path string, vp *Viewport, xi *rpc.Connection) (*View, error) {
	view := &View{}
	view.view = NewViewport(vp, 3, 0)
	view.gutter = NewViewport(vp, 0, 0)
	view.xi = xi
	view.LineCache = NewLineCache()

	msg, err := xi.Request(&rpc.Request{
		Method: "new_view",
		Params: &rpc.Object{"file_path": path},
	})
	if err != nil {
		return view, err
	}

	view.ID = msg.Value.(string)
	view.InputHandler = &InputHandler{view.ID, path, xi}

	// Set scroll window size
	_, height := vp.Size()
	xi.Notify(&rpc.Request{
		Method: "edit",
		Params: &rpc.Object{
			"method":  "scroll",
			"params":  &rpc.Array{0, height - 2},
			"view_id": view.ID,
		},
	})

	return view, nil
}

func (v *View) Draw() {
	if len(v.lines) == 0 {
		return
	}

	// render line numbers
	// TODO: improve, alot... :-)
	style := defaultStyle.Foreground(tcell.ColorLightCyan)
	width := len(strconv.Itoa(len(v.lines) + v.LineCache.invalidBefore))

	v.gutter.SetWidth(width + 1)
	v.view.SetOffsetX(width + 2)
	for i := 0; i < len(v.lines); i++ {
		nLine := i + v.invalidBefore
		txt := ralign(strconv.Itoa(nLine+1), width)
		width := len(txt)
		for x := 0; x < width; x++ {
			v.gutter.SetContent(1+x, nLine, rune(txt[x]), nil, style)
		}
	}

	// TODO: Line numbers in a separate viewport
	for y, line := range v.lines {
		if line == nil {
			continue
		}
		nLine := y + v.invalidBefore
		visualX := 0
		for x, char := range []rune(line.Text) {
			// TODO: Do this somewhere else
			var style = defaultStyle
			// TODO: Reserved??
			if line.StyleIds[x] >= 2 {
				fg, _, _ := styles[line.StyleIds[x]].Decompose()
				style = style.Foreground(fg)
			}

			if char == '\t' {
				ts := tabSize - (visualX % tabSize)
				for i := 0; i < ts; i++ {
					v.view.SetContent(visualX+i, nLine, ' ', nil, style)
				}
				visualX += ts
			} else if char != '\n' {
				// TODO: Trim newline in a better way?
				v.view.SetContent(visualX, nLine, char, nil, style)
				visualX++
			}
		}
		if len(line.Cursors) != 0 {
			// TODO: Verify if xi-core will take care of tabs for us
			cX := GetCursorVisualX(line.Cursors[0], line.Text)
			// TODO: Multiple cursor support
			v.view.ShowCursor(cX, nLine)
		}
	}
	lineStart, lineEnd := v.view.GetViewport()
	if lineStart != v.lineStart || lineEnd != v.lineEnd {
		v.Scroll(lineStart, lineEnd)
		v.lineStart = lineStart
		v.lineEnd = lineEnd
	}
}

func (v *View) MakeLineVisible(line int) {
	v.view.MakeVisibleY(line)
	v.gutter.MakeVisibleY(line)
}

// HandleEvent handles tcell events
func (v *View) HandleEvent(ev tcell.Event) {
	switch e := ev.(type) {
	case *tcell.EventMouse:
		x, y := e.Position()
		log.Println(x, y, e.Buttons(), e.Modifiers())
		switch e.Buttons() {
		case tcell.Button1:
			v.Click(y, x-v.gutter.width-1, 0, 1)
		}
	case *tcell.EventKey:
		ctrl := e.Modifiers()&tcell.ModCtrl != 0
		alt := e.Modifiers()&tcell.ModAlt != 0
		log.Println(string(e.Name()))
		if e.Key() == tcell.KeyRune && !ctrl && !alt {
			v.Insert(string(e.Rune()))
		} else {
			if ctrl && alt {
				switch e.Name() {
				case "Alt+Ctrl+L":
					//	log.Println(goPlugin.Format(path))

				default:
					log.Println(string(e.Name()))
				}
			} else if ctrl && !alt {
				switch e.Key() {
				case tcell.KeyLeft:
					v.MoveWordLeft()
				case tcell.KeyRight:
					v.MoveWordRight()
				case tcell.KeyCtrlS:
					v.Save()
				case tcell.KeyCtrlZ:
					v.Undo()
				case tcell.KeyCtrlR:
					v.Redo()
				case tcell.KeyCtrlL:

				case tcell.KeyCtrlD:
					v.DuplicateLine()
				case tcell.KeyCtrlV:
					s, e := clipboard.ReadAll()
					if e != nil {
						log.Println(e)
						break
					}
					v.Insert(s)
				default:
					log.Println(string(e.Name()))
				}
			} else {
				switch e.Key() {
				case tcell.KeyBackspace2, tcell.KeyBackspace:
					v.DeleteBackward()
				case tcell.KeyTAB:
					v.Tab()
				case tcell.KeyEnter:
					v.Newline()
				case tcell.KeyLeft:
					v.MoveLeft()
				case tcell.KeyUp:
					v.MoveUp()
				case tcell.KeyRight:
					v.MoveRight()
				case tcell.KeyDown:
					v.MoveDown()
				case tcell.KeyDelete:
					v.DeleteForward()
				default:
					log.Println(string(e.Name()))
				}
			}
		}
	}
}
