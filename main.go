package main

import (
	"flag"
	"github.com/andlabs/ui"
)

var buffer = Buffer{}

func main() {
	gui := flag.Bool("w", false, "window gui")
	flag.Parse()

	if len(flag.Args()) == 1 {
		buffer.Open(flag.Args()[0])
	}
	if len(flag.Args()) == 0 {
		buffer.OpenNew()
	}

	if *gui {
		ui.Main(setupUI)
	} else {
		cli()
	}
}
