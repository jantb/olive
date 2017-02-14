package main

import (
	"bufio"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/jantb/rope"
)

//Buffer holds the current buffer
type Buffer struct {
	r        *rope.RopeRuneRope
	filename string
}

// Open initializes a new buffer
func (b *Buffer) Open(filename string) {
	file, err := os.Open(filename)
	b.filename = filename
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	b.r = rope.NewFromRuneRope([]rope.RuneRope{})

	ropes := make([]rope.RuneRope, 0, 512)

	i := 0
	for {
		line, _, err := reader.ReadLine()

		if err == io.EOF {
			break
		}

		ropes = append(ropes, *rope.NewFromRunes([]rune(string(line) + "\n")))
		if i == 4096 {
			b.r = b.r.Insert(b.r.Len(), ropes)
			ropes = make([]rope.RuneRope, 0, 512)
			i = 0
			continue
		}
		i++
	}
	b.r = b.r.Insert(b.r.Len(), ropes)
}

// GetLines return lines from the buffer
func (b *Buffer) GetLines(top, left, length, w int) [][]rune {
	ret := make([][]rune, length)

	for i, rope := range b.r.Sub(top, length) {
		if rope.Len() < left {
			ret[i] = []rune{}
			continue
		}
		ret[i] = rope.Sub(left, w)
	}
	return ret
}

// GetLine return line from the buffer
func (b *Buffer) GetLine(line int) []rune {
	if buffer.Len() > line {
		rr := b.r.Index(line)
		return rr.Runes()
	}
	return []rune{}
}

// GetLineLen return line len from the buffer
func (b *Buffer) GetLineLen(line int) int {
	if buffer.Len() > line {
		rr := b.r.Index(line)
		return rr.Len()
	}
	return 0
}

// Insert into the buffer
func (b *Buffer) Insert(row, column int, bytes []rune) {
	// Empthy buffer just insert a new line
	if b.r.Len() <= row {
		b.r = b.r.Insert(0, []rope.RuneRope{*rope.NewFromRunes([]rune(""))})
	}
	r := b.r.Index(row)
	b.r = b.r.Delete(row, 1)
	b.r = b.r.Insert(row, []rope.RuneRope{*r.Insert(column, bytes)})
}

// InsertEnter into the buffer
func (b *Buffer) InsertEnter(row, column int) {
	if b.r.Len() < row+1 {
		b.r = b.r.Insert(row, []rope.RuneRope{*rope.NewFromRunes([]rune("\n"))})
		return
	}
	ro := b.r.Index(row)
	if column == ro.Len()-1 {
		b.r = b.r.Insert(row+1, []rope.RuneRope{*rope.NewFromRunes([]rune("\n"))})
		return
	}
	l, r := ro.Split(column)
	b.r = b.r.Delete(row, 1)
	l = l.Insert(l.Len(), []rune{'\n'})
	if r == nil {
		b.r = b.r.Insert(row, []rope.RuneRope{*l})
	} else {
		b.r = b.r.Insert(row, []rope.RuneRope{*l, *r})
	}
}

func (b *Buffer) Backspace(row, column int) int {
	// if row == 0 && column == 0 {
	// 	return 0
	// }
	// if column == 0 {
	// 	b.Delete(row-1, len(b.GetLine(row-1)))
	// 	return 0
	// }
	return b.Delete(row, column-1)
}

// Delete char from
func (b *Buffer) Delete(row, column int) int {
	if column == -1 {
		if row > 0 {
			prevr := b.r.Index(row - 1)
			prevr = *prevr.Delete(prevr.Len()-1, 1)
			if b.Len() > row {
				r := b.r.Index(row)
				if r.Len() > 0 {
					prevr = *prevr.Concat(&r)
				}
				b.r = b.r.Delete(row-1, 2)
				b.r = b.r.Insert(row, []rope.RuneRope{prevr})
				return r.Len()
			} else {
				r := b.r.Index(row - 1)
				if r.Len() > 1 {
					prevr = *prevr.Concat(&r)
				}
				b.r = b.r.Delete(row-1, 2)
				b.r = b.r.Insert(row-1, []rope.RuneRope{prevr})
				return r.Len()
			}
		} else {
			buffer.RemoveRow(0)
		}
	} else {
		r := b.r.Index(row)
		if row == 0 && column == 0 && r.Len() == 1 {
			buffer.RemoveRow(0)
			return 0
		}
		b.r = b.r.Delete(row, 1)
		b.r = b.r.Insert(row, []rope.RuneRope{*r.Delete(column, 1)})
		return 0
	}
	return 0
}

// RemoveRow removes row
func (b *Buffer) RemoveRow(row int) {
	b.r = b.r.Delete(row, 1)
}

// Len gives length of buffer
func (b *Buffer) Len() int {
	return b.r.Len()
}

// New buffer
func (b *Buffer) New() {
	b.r = rope.NewFromRuneRope([]rope.RuneRope{})
}

// Bytes returns all bytes from buffer
func (b *Buffer) Bytes() []byte {
	return b.r.Bytes()
}

// Save the buffer
func (b *Buffer) Save() {
	ioutil.WriteFile(b.filename, b.Bytes(), 0644)
}
