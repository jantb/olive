package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jantb/rope"
)

//Buffer holds the current buffer
type Buffer struct {
	r *rope.RopeRope
}

// Open initializes a new buffer
func (b *Buffer) Open(filename string) {
	file, err := os.Open("big.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	b.r = rope.NewFromRope([]rope.Rope{})
	t := time.Now()

	ropes := []rope.Rope{}

	for scanner.Scan() {
		ropes = append(ropes, *rope.NewFromBytes([]byte(scanner.Text() + "\n")))
	}
	b.r = b.r.Insert(b.r.Len(), ropes)

	fmt.Print(time.Now().Sub(t))
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func (b *Buffer) getLines(start, length int) [][]byte {
	ret := make([][]byte, length)
	for i, rope := range b.r.Sub(start, length) {
		ret[i] = rope.Bytes()
	}
	return ret
}
