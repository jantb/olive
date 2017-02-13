package main

//Loc holds a location
type Loc struct {
	row    int
	column int
}

// Cursor object
type Cursor struct {
	loc *Loc
}

type cursors []Cursor

var c = cursors{}

// SetCursor removes all other cursor and sets current
func SetCursor(row, column int) {
	c = c[:0]
	c = append(c, Cursor{loc: &Loc{row, column}})
}

// GetCursor get first cursor
func GetCursor() Cursor {
	if len(c) == 0 {
		SetCursor(0, 0)
	}
	return c[0]
}

// MoveRight the cursor
func (c *Cursor) MoveRight() {
	if c.loc.column < len(buffer.GetLine(c.loc.row)) {
		c.loc.column++
		c.showCursorInView()
	}
}

// MoveLeft the cursor
func (c *Cursor) MoveLeft() {
	if c.loc.column > 0 {
		c.loc.column--
		c.showCursorInView()
	}
}

// MoveDown the cursor
func (c *Cursor) MoveDown() {
	if buffer.Len() >= c.loc.row+1 {
		c.loc.row++
		for c.loc.column > len(buffer.GetLine(c.loc.row)) {
			c.MoveLeft()
		}
		c.showCursorInView()
	}
}

// MovePageDown the cursor
func (c *Cursor) MovePageDown() {
	if buffer.Len() > c.loc.row+height {
		c.loc.row += height
	} else {
		c.loc.row = buffer.Len()
	}
	for c.loc.column > len(buffer.GetLine(c.loc.row)) {
		c.MoveLeft()
	}
	c.showCursorInView()
}

// MovePageUp the cursor
func (c *Cursor) MovePageUp() {
	if c.loc.row-height >= 0 {
		c.loc.row -= height
	} else {
		c.loc.row = 0
	}
	for c.loc.column > len(buffer.GetLine(c.loc.row)) {
		c.MoveLeft()
	}
	c.showCursorInView()
}

// MoveStartOfLine the cursor to start of line
func (c *Cursor) MoveStartOfLine() {
	c.loc.column = 0
	c.showCursorInView()
}

// MoveToEndOfLine the cursor to end of line
func (c *Cursor) MoveToEndOfLine() {
	c.loc.column = len(buffer.GetLine(c.loc.row))
	c.showCursorInView()
}

// MoveUp the cursor
func (c *Cursor) MoveUp() {
	if c.loc.row > 0 {
		c.loc.row--
		for c.loc.column > len(buffer.GetLine(c.loc.row)) {
			c.MoveLeft()
		}
		c.showCursorInView()
	}
}

func (c Cursor) showCursorInView() {
	for c.loc.row > topRow+height-6 {
		topRow = Min(buffer.Len()-5, topRow+1)
	}
	for c.loc.row < topRow+5 && topRow > 0 {
		topRow = Max(0, topRow-1)
	}
	for c.loc.column < leftColumn+5 && leftColumn > 0 {
		leftColumn = Max(0, leftColumn-1)
	}
	for c.loc.column > leftColumn+width-offset-5 {
		lengthOfLine := len(buffer.GetLine(c.loc.row))
		leftColumn = Min(lengthOfLine, leftColumn+1)
	}
}
