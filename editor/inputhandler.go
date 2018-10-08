package editor

import (
	"github.com/jantb/olive/rpc"
)

type InputHandler struct {
	ViewID   string
	FilePath string
	c        *rpc.Connection
}

func (ih *InputHandler) edit(params rpc.Object) {
	params["view_id"] = ih.ViewID
	if _, ok := params["params"]; !ok {
		params["params"] = &rpc.Object{}
	}

	ih.c.Notify(&rpc.Request{
		Method: "edit",
		ViewID: ih.ViewID,
		Params: params,
	})
}

func (ih *InputHandler) MoveLeft() {
	ih.edit(rpc.Object{"method": "move_left"})
}

func (ih *InputHandler) MoveRight() {
	ih.edit(rpc.Object{"method": "move_right"})
}

func (ih *InputHandler) MoveUp() {
	ih.edit(rpc.Object{"method": "move_up"})
}

func (ih *InputHandler) MoveDown() {
	ih.edit(rpc.Object{"method": "move_down"})
}

func (ih *InputHandler) MoveWordLeft() {
	ih.edit(rpc.Object{"method": "move_word_left"})
}

func (ih *InputHandler) MoveWordRight() {
	ih.edit(rpc.Object{"method": "move_word_right"})
}

func (ih *InputHandler) DeleteBackward() {
	ih.edit(rpc.Object{"method": "delete_backward"})
}

func (ih *InputHandler) DeleteForward() {
	ih.edit(rpc.Object{"method": "delete_forward"})
}

func (ih *InputHandler) Undo() {
	ih.edit(rpc.Object{"method": "undo"})
}

func (ih *InputHandler) Redo() {
	ih.edit(rpc.Object{"method": "redo"})
}
func (ih *InputHandler) GoToLine(line int) {
	ih.edit(rpc.Object{"method": "goto_line", "params": rpc.Object{"line": line}})
}

func (ih *InputHandler) DuplicateLine() {
	ih.edit(rpc.Object{"method": "duplicate_line"})
}

func (ih *InputHandler) Tab() {
	ih.edit(rpc.Object{"method": "insert_tab"})
}

func (ih *InputHandler) Newline() {
	ih.edit(rpc.Object{"method": "insert_newline"})
}

func (ih *InputHandler) Insert(char string) {
	ih.c.Notify(&rpc.Request{
		Method: "edit",
		ViewID: ih.ViewID,
		Params: rpc.Object{"method": "insert", "params": rpc.Object{"chars": char}, "view_id": ih.ViewID},
	})
}

func (ih *InputHandler) Scroll(lineStart, lineEnd int) {
	ih.edit(rpc.Object{"method": "scroll", "params": rpc.Array{lineStart, lineEnd}})
}

func (ih *InputHandler) Save() {
	ih.c.Notify(&rpc.Request{
		Method: "save",
		Params: rpc.Object{
			"view_id":   ih.ViewID,
			"file_path": ih.FilePath,
		},
	})
}

func (ih *InputHandler) Click(x, y, mod, clicks int) {
	ih.edit(rpc.Object{"method": "click", "params": rpc.Array{y, x, mod, clicks}})
}
