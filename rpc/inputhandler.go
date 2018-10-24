package rpc

import (
	"log"
)

type InputHandler struct {
	ViewID   string
	FilePath string
	Xi       *Connection
}

func (ih *InputHandler) edit(params Object) {
	params["view_id"] = ih.ViewID
	if _, ok := params["params"]; !ok {
		params["params"] = &Object{}
	}

	ih.Xi.Notify(&Request{
		Method: "edit",
		Params: params,
	})
}
func (ih *InputHandler) gesture(params Object) {
	params["view_id"] = ih.ViewID
	if _, ok := params["params"]; !ok {
		params["params"] = &Object{}
	}

	ih.Xi.Notify(&Request{
		Method: "gesture",
		Params: params,
	})
}
func (ih *InputHandler) MoveLeft() {
	ih.edit(Object{"method": "move_left"})
}

func (ih *InputHandler) MoveLeftAndModifySelection() {
	ih.edit(Object{"method": "move_left_and_modify_selection"})
}

func (ih *InputHandler) SelectAll() {
	ih.edit(Object{"method": "select_all"})
}

func (ih *InputHandler) AddSelectionAbove() {
	ih.edit(Object{"method": "add_selection_above"})
}

func (ih *InputHandler) AddSelectionBelow() {
	ih.edit(Object{"method": "add_selection_below"})
}
func (ih *InputHandler) CancelOperation() {
	ih.edit(Object{"method": "cancel_operation"})
}

func (ih *InputHandler) MoveRight() {
	ih.edit(Object{"method": "move_right"})
}

func (ih *InputHandler) MoveRightAndModifySelection() {
	ih.edit(Object{"method": "move_right_and_modify_selection"})
}

func (ih *InputHandler) MoveUp() {
	ih.edit(Object{"method": "move_up"})
}

func (ih *InputHandler) MoveUpAndModifySelection() {
	ih.edit(Object{"method": "move_up_and_modify_selection"})
}

func (ih *InputHandler) MoveDown() {
	ih.edit(Object{"method": "move_down"})
}

func (ih *InputHandler) MoveDownAndModifySelection() {
	ih.edit(Object{"method": "move_down_and_modify_selection"})
}

func (ih *InputHandler) ScrollPageUp() {
	ih.edit(Object{"method": "scroll_page_up"})
}

func (ih *InputHandler) ScrollPageDown() {
	ih.edit(Object{"method": "scroll_page_down"})
}

func (ih *InputHandler) ScrollPageUpAndModifySelection() {
	ih.edit(Object{"method": "scroll_page_up_and_modify_selection"})
}

func (ih *InputHandler) ScrollPageDownAndModifySelection() {
	ih.edit(Object{"method": "scroll_page_down_and_modify_selection"})
}

func (ih *InputHandler) MoveToBeginningOfLine() {
	ih.edit(Object{"method": "move_to_left_end_of_line"})
}

func (ih *InputHandler) MoveToBeginningOfLineAndModifySelection() {
	ih.edit(Object{"method": "move_to_left_end_of_line_and_modify_selection"})
}

func (ih *InputHandler) MoveToEndOfLine() {
	ih.edit(Object{"method": "move_to_right_end_of_line"})
}

func (ih *InputHandler) MoveToEndOfLineAndModifySelection() {
	ih.edit(Object{"method": "move_to_right_end_of_line_and_modify_selection"})
}

func (ih *InputHandler) MoveWordLeft() {
	ih.edit(Object{"method": "move_word_left"})
}

func (ih *InputHandler) MoveWordLeftAndModifySelection() {
	ih.edit(Object{"method": "move_word_left_and_modify_selection"})
}

func (ih *InputHandler) MoveWordRightAndModifySelection() {
	ih.edit(Object{"method": "move_word_right_and_modify_selection"})
}

func (ih *InputHandler) MoveWordRight() {
	ih.edit(Object{"method": "move_word_right"})
}

func (ih *InputHandler) DeleteBackward() {
	ih.edit(Object{"method": "delete_backward"})
}

func (ih *InputHandler) DeleteWordForward() {
	ih.edit(Object{"method": "delete_word_forward"})
}

func (ih *InputHandler) DeleteLine() {
	ih.SelectLine()
	ih.DeleteForward()
}

func (ih *InputHandler) MoveToBeginningOfDocument() {
	ih.edit(Object{"method": "move_to_beginning_of_document"})
}

func (ih *InputHandler) MoveToEndOfDocument() {
	ih.edit(Object{"method": "move_to_end_of_document"})
}

func (ih *InputHandler) DeleteWordBackward() {
	ih.edit(Object{"method": "delete_word_backward"})
}

func (ih *InputHandler) DeleteForward() {
	ih.edit(Object{"method": "delete_forward"})
}
func (ih *InputHandler) SelectLine() {
	ih.MoveToBeginningOfLine()
	ih.MoveToEndOfLineAndModifySelection()
	ih.MoveRightAndModifySelection()
}

func (ih *InputHandler) Undo() {
	ih.edit(Object{"method": "undo"})
}

func (ih *InputHandler) Redo() {
	ih.edit(Object{"method": "redo"})
}

func (ih *InputHandler) Transpose() {
	ih.edit(Object{"method": "transpose"})
}

func (ih *InputHandler) Yank() {
	ih.edit(Object{"method": "yank"})
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

func (ih *InputHandler) Find(chars string, caseSensitive, regex, wholeWords bool) {
	ih.edit(Object{"method": "find", "params": Object{
		"chars":          chars,
		"case_sensitive": false,
		"regex":          false,
		"whole_words":    false,
	}})
}

// find_next {"wrap_around": true, "allow_same": false, "modify_selection": "set"} find_previous {"wrap_around": true, "allow_same": false, "modify_selection": "set"} All parameters are optional. Boolean parameters are by default false and modify_selection is set by default. If allow_same is set to true the current selection is considered a valid next occurrence. Supported options for modify_selection are:
//none: the selection is not modified
//set: the next/previous match will be set as the new selection
//add: the next/previous match will be added to the current selection
//add_remove_current: the previously added selection will be removed and the next/previous match will be added to the current selection
//Selects the next/previous occurrence matching the search query.
func (ih *InputHandler) FindNext(wrapAround, allowSame bool, modifySelection string) {
	ih.edit(Object{"method": "find_next", "params": Object{
		"wrap_around": wrapAround, "allow_same": allowSame, "modify_selection": modifySelection,
	}})
}
func (ih *InputHandler) Click(x, y, mod, clicks int) {
	ih.edit(Object{"method": "click", "params": Array{y, x, mod, clicks}})
}

func (ih *InputHandler) RequestLines(first, last int) {
	ih.edit(Object{"method": "request_lines", "params": Array{first, last}})
}

func (ih *InputHandler) MoveLineUp() {
	ih.SelectLine()
	s := ih.Cut()
	ih.MoveUp()
	ih.SelectLine()
	s2 := ih.Cut()
	ih.Insert(s)
	ih.MoveDown()
	ih.Insert(s2)
	ih.MoveUp()
}

func (ih *InputHandler) MoveLineDown() {
	ih.SelectLine()
	s := ih.Cut()
	ih.MoveDown()
	ih.SelectLine()
	s2 := ih.Cut()
	ih.Insert(s)
	ih.MoveUp()
	ih.Insert(s2)
	ih.MoveDown()
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

func (ih *InputHandler) Copy() string {

	msg, err := ih.Xi.Request(&Request{
		Method: "edit",
		Params: Object{"method": "copy", "view_id": ih.ViewID},
	})
	if err != nil {
		log.Fatal(err)
	}

	if msg.Value == nil {
		return ""
	}
	return msg.Value.(string)
}

func (ih *InputHandler) Cut() string {

	msg, err := ih.Xi.Request(&Request{
		Method: "edit",
		Params: Object{"method": "cut", "view_id": ih.ViewID},
	})

	if err != nil {
		log.Fatal(err)
	}
	if msg.Value == nil {
		return ""
	}
	return msg.Value.(string)
}
