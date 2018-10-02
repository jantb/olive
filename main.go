package main

import (
	"flag"
	"github.com/andlabs/ui"
)

var buffer = Buffer{}

func main() {
	gui := flag.Bool("w", false, "window gui")
	flag.Parse()
	if *gui {
		ui.Main(setupUI)
	} else {
		cli()
	}
}
