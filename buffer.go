package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/jantb/rope"
)

//Buffer holds the current buffer
type Buffer struct {
	r *rope.RopeRuneRope
}

// Open initializes a new buffer
func (b *Buffer) Open(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	b.r = rope.NewFromRuneRope([]rope.RuneRope{})

	ropes := make([]rope.RuneRope, 0, 512)

	t := time.Now()
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

	fmt.Print(time.Now().Sub(t))

}

// GetLines return lines from the buffer
func (b *Buffer) GetLines(start, length, w int) [][]rune {
	ret := make([][]rune, length)
	for i, rope := range b.r.Sub(start, length) {
		ret[i] = rope.Sub(0, w)
	}
	return ret
}

// Insert into the buffer
func (b *Buffer) Insert(row, column int, bytes []rune) {
	for row >= b.r.Len() {
		b.r = b.r.Insert(b.r.Len(), []rope.RuneRope{*rope.NewFromRunes([]rune(""))})
	}
	r := b.r.Index(row)
	for r.Len() < column {
		r = *r.Insert(column, []rune(" "))
	}
	b.r = b.r.Delete(row, 1)

	b.r = b.r.Insert(row, []rope.RuneRope{*r.Insert(column, bytes)})
}

// Delete char from
func (b *Buffer) Delete(row, column int) {
	for row >= b.r.Len() {
		b.r = b.r.Insert(b.r.Len(), []rope.RuneRope{*rope.NewFromRunes([]rune(""))})
	}
	r := b.r.Index(row)
	for r.Len() < column {
		r = *r.Insert(column, []rune(" "))
	}
	if column < 0 {
		return
	}
	b.r = b.r.Delete(row, 1)
	b.r = b.r.Insert(row, []rope.RuneRope{*r.Delete(column, 1)})
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
