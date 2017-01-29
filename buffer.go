package main

import (
	"io/ioutil"
	"log"

	"github.com/jantb/rope"
)

//Buffer holds the current buffer
type Buffer struct {
	r *rope.Rope
}

// Open initializes a new buffer
func (b *Buffer) Open(filename string) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Panic(err)
	}
	b.r = rope.NewFromBytes(bytes)
}
