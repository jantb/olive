package main

var buffer = Buffer{}

func main() {
	buffer.Open("main.go")
	Display(&buffer)
}
