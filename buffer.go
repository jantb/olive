package main

import (
	"bufio"
	"log"
	"os"

	"github.com/jantb/rope"
)

//Buffer holds the current buffer
type Buffer struct {
	r  *rope.Rope
	rr *rope.RopeRope
}

// Open initializes a new buffer
func (b *Buffer) Open(filename string) {
	file, err := os.Open("big.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	b.rr = rope.NewFromRope([]rope.Rope{})
	for scanner.Scan() {
		by := []byte(scanner.Text())
		by = append(by, []byte("\n")...)
		r := rope.NewFromBytes(by)
		b.rr = b.rr.Insert(b.rr.Len(), []rope.Rope{*r})
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func (b *Buffer) getLines(start, length int) [][]byte {
	ret := make([][]byte, length)
	for i, rope := range b.rr.Sub(start, length) {
		ret[i] = rope.Bytes()
	}
	return ret
}
