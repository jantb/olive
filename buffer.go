package main

import (
	"io/ioutil"
	"log"
	"strings"
)

//Buffer holds the current buffer
type Buffer struct {
	r        *Rope
	filename string
}

// Open initializes a new buffer
func (b *Buffer) Open(filename string) {
	bytes, err := ioutil.ReadFile(filename)
	b.filename = filename
	if err != nil {
		log.Fatal(err)
	}

	b.r = NewFromRunes([]rune(string(bytes)))
}

// GetLines return lines from the buffer
func (b *Buffer) GetLines(top, left, length, w int) [][]rune {
	ret := make([][]rune, length)
	for key, value := range b.r.GetLines(top, top+length) {
		ret[key] = value
	}

	return ret
}

// GetLine return line from the buffer
func (b *Buffer) GetLine(line int) []rune {
	return b.r.GetLine(line)
}

// GetLineLen return line len from the buffer
func (b *Buffer) GetLineLen(line int) int {
	return len(b.r.GetLine(line)) - 1
}

// Insert into the buffer
func (b *Buffer) Insert(row, column int, runes []rune) {
	b.r = b.r.Insert(b.r.GetLineOffset(0, row)+column, runes)
}

// InsertEnter into the buffer
func (b *Buffer) InsertEnter(row, column int) {
	b.Insert(row, column, []rune{'\n'})
}

func (b *Buffer) Backspace(row, column int) int {
	return b.Delete(row, column)
}

// Delete char from
func (b *Buffer) Delete(row, column int) int {
	b.r = b.r.Delete(b.r.GetLineOffset(0, row)+column-1, 1)
	return -1
}

// RemoveRow removes row
func (b *Buffer) RemoveRow(row int) {
	lfs := b.r.GetLineOffset(0, row)
	lfe := b.r.GetLineOffset(0, row+1)
	b.r = b.r.Delete(lfs, lfe-lfs)
}

// Len gives length of buffer
func (b *Buffer) Len() int {
	return len(strings.Split(string(b.r.Runes()), "\n"))
}

// New buffer
func (b *Buffer) New() {
	b.r = NewFromRunes([]rune(""))
}

// NewText buffer From Text
func (b *Buffer) NewText(text string) {
	b.r = NewFromRunes([]rune(text))
}

// Bytes returns all bytes from buffer
func (b *Buffer) Bytes() []byte {
	return []byte(string(b.r.Runes()))
}

// Save the buffer
func (b *Buffer) Save() {
	ioutil.WriteFile(b.filename, b.Bytes(), 0644)
}
