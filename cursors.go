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
	c.loc.column++
}

// MoveLeft the cursor
func (c *Cursor) MoveLeft() {
	if c.loc.column > 0 {
		c.loc.column--
	}
}

// MoveDown the cursor
func (c *Cursor) MoveDown() {
	c.loc.row++
}

// MoveUp the cursor
func (c *Cursor) MoveUp() {
	if c.loc.row > 0 {
		c.loc.row--
	}
}
