package main

import (
	"log"
	"os"
	"os/user"
	"path/filepath"
	"syscall"
)

var buffer = Buffer{}

func main() {

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	logFile, _ := os.OpenFile(filepath.Join(usr.HomeDir, ".olive.log"), os.O_WRONLY|os.O_CREATE|os.O_SYNC, 0755)
	syscall.Dup2(int(logFile.Fd()), 1)
	syscall.Dup2(int(logFile.Fd()), 2)

	if len(os.Args) == 2 {
		buffer.Open(os.Args[1])
		loadDark()
		transforSyntax()
		Display(&buffer)
	}
}
