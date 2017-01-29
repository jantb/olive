package main

var buffer = Buffer{}

func main() {
	buffer.Open("display.go")
	Display(&buffer)
}
