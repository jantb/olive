package rpc

import "github.com/linde12/kod/rpc"

type InputHandler struct {
	ViewID   string
	FilePath string
	C        *Connection
}

func (ih *InputHandler) edit(params Object) {
	params["view_id"] = ih.ViewID
	if _, ok := params["params"]; !ok {
		params["params"] = &rpc.Object{}
	}

	ih.C.Notify(&Request{
		Method: "edit",
		ViewID: ih.ViewID,
		Params: params,
	})
}

func (ih *InputHandler) MoveLeft() {
	ih.edit(Object{"method": "move_left"})
}

func (ih *InputHandler) MoveRight() {
	ih.edit(Object{"method": "move_right"})
}

func (ih *InputHandler) MoveUp() {
	ih.edit(Object{"method": "move_up"})
}

func (ih *InputHandler) MoveDown() {
	ih.edit(Object{"method": "move_down"})
}

func (ih *InputHandler) MoveWordLeft() {
	ih.edit(Object{"method": "move_word_left"})
}

func (ih *InputHandler) MoveWordRight() {
	ih.edit(Object{"method": "move_word_right"})
}

func (ih *InputHandler) DeleteBackward() {
	ih.edit(Object{"method": "delete_backward"})
}

func (ih *InputHandler) DeleteForward() {
	ih.edit(Object{"method": "delete_forward"})
}

func (ih *InputHandler) Undo() {
	ih.edit(Object{"method": "undo"})
}

func (ih *InputHandler) Redo() {
	ih.edit(Object{"method": "redo"})
}
func (ih *InputHandler) GoToLine(line int) {
	ih.edit(Object{"method": "goto_line", "params": Object{"line": line}})
}

func (ih *InputHandler) DuplicateLine() {
	ih.edit(Object{"method": "duplicate_line"})
}

func (ih *InputHandler) Tab() {
	ih.edit(Object{"method": "insert_tab"})
}

func (ih *InputHandler) Newline() {
	ih.edit(Object{"method": "insert_newline"})
}

func (ih *InputHandler) Insert(char string) {
	ih.C.Notify(&Request{
		Method: "edit",
		ViewID: ih.ViewID,
		Params: Object{"method": "insert", "params": Object{"chars": char}, "view_id": ih.ViewID},
	})
}

func (ih *InputHandler) Scroll(lineStart, lineEnd int) {
	ih.edit(Object{"method": "scroll", "params": Array{lineStart, lineEnd}})
}

func (ih *InputHandler) Save() {
	ih.C.Notify(&Request{
		Method: "save",
		Params: Object{
			"view_id":   ih.ViewID,
			"file_path": ih.FilePath,
		},
	})
}

func (ih *InputHandler) Click(x, y, mod, clicks int) {
	ih.edit(Object{"method": "click", "params": Array{y, x, mod, clicks}})
}

func (ih *InputHandler) RequestLines(first, last int) {

}

func (ih *InputHandler) MoveLineUp() {

}

func (ih *InputHandler) MoveLineDown() {

}
