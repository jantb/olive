package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sync"
	"time"

	"github.com/jantb/rope"
)

//Buffer holds the current buffer
type Buffer struct {
	r *rope.RopeRope
}

var m sync.Mutex

// Open initializes a new buffer
func (b *Buffer) Open(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	b.r = rope.NewFromRope([]rope.Rope{})

	ropes := make([]rope.Rope, 0, 512)

	t := time.Now()
	i := 0
	for {
		line, _, err := reader.ReadLine()

		if err == io.EOF {
			break
		}

		ropes = append(ropes, *rope.NewFromBytes([]byte(string(line) + "\n")))
		if i == 512 {
			go func(ropes *[]rope.Rope) {
				m.Lock()
				b.r = b.r.Insert(b.r.Len(), *ropes)
				m.Unlock()
			}(&ropes)
			ropes = make([]rope.Rope, 0, 512)
			i = 0
			continue
		}
		i++
	}
	b.r = b.r.Insert(b.r.Len(), ropes)

	fmt.Print(time.Now().Sub(t))

}

// GetLines return lines from the buffer
func (b *Buffer) GetLines(start, length, w int) [][]byte {
	ret := make([][]byte, length)
	for i, rope := range b.r.Sub(start, length) {
		ret[i] = rope.Sub(0, w)
	}
	return ret
}

// Insert into the buffer
func (b *Buffer) Insert(row, column int, bytes []byte) {
	for row >= b.r.Len() {
		b.r = b.r.Insert(b.r.Len(), []rope.Rope{*rope.NewFromBytes([]byte("\n"))})
	}
	r := b.r.Index(row)
	b.r = b.r.Delete(row, 1)
	b.r = b.r.Insert(row, []rope.Rope{*r.Insert(column, bytes)})
}

// Insert into the buffer
func (b *Buffer) RemoveRow(row int) {
	b.r = b.r.Delete(row, 1)
}

// Len gives length of buffer
func (b *Buffer) Len() int {
	return b.r.Len()
}

// New buffer
func (b *Buffer) New() {
	b.r = rope.NewFromRope([]rope.Rope{})
}
