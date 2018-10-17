package rpc

import (
	"github.com/linde12/kod/rpc"
)

type InputHandler struct {
	ViewID   string
	FilePath string
	Xi       *Connection
}

func (ih *InputHandler) edit(params Object) {
	params["view_id"] = ih.ViewID
	if _, ok := params["params"]; !ok {
		params["params"] = &rpc.Object{}
	}

	ih.Xi.Notify(&Request{
		Method: "edit",
		Params: params,
	})
}
func (ih *InputHandler) gesture(params Object) {
	params["view_id"] = ih.ViewID
	if _, ok := params["params"]; !ok {
		params["params"] = &rpc.Object{}
	}

	ih.Xi.Notify(&Request{
		Method: "gesture",
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

func (ih *InputHandler) ScrollPageUp() {
	ih.edit(Object{"method": "scroll_page_up"})
}

func (ih *InputHandler) ScrollPageDown() {
	ih.edit(Object{"method": "scroll_page_down"})
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
	ih.Xi.Notify(&Request{
		Method: "edit",
		Params: Object{"method": "insert", "params": Object{"chars": char}, "view_id": ih.ViewID},
	})
}

func (ih *InputHandler) Scroll(lineStart, lineEnd int) {
	ih.edit(Object{"method": "scroll", "params": Array{lineStart, lineEnd}})
}

func (ih *InputHandler) Resize(width, height int) {
	ih.edit(Object{"method": "resize", "params": Object{"width": width, "height": height}})
}

func (ih *InputHandler) Save() {
	ih.Xi.Notify(&Request{
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
	ih.edit(Object{"method": "request_lines", "params": Array{first, last}})
}

func (ih *InputHandler) MoveLineUp() {
	//RequestLines//
	//DeleteToBeginningOfLine
	//MoveToRightEndOfLine

}

func (ih *InputHandler) MoveLineDown() {

}
func (ih *InputHandler) Close() {
	ih.Xi.Notify(&Request{
		Method: "close_view",
		Params: Object{
			"view_id": ih.ViewID,
		},
	})
}

func (ih *InputHandler) AddCursor(x, y int) {
	ih.edit(Object{"method": "gesture", "params": Object{"line": y, "col": x, "ty": "toggle_sel"}})
}
