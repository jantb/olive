package main

import (
	"log"
	"os"
	"os/user"
	"path/filepath"
	"syscall"
)

func cli() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	os.Remove(filepath.Join(usr.HomeDir, ".olive.log"))
	logFile, _ := os.OpenFile(filepath.Join(usr.HomeDir, ".olive.log"), os.O_WRONLY|os.O_CREATE|os.O_SYNC, 0755)
	syscall.Dup2(int(logFile.Fd()), 1)
	syscall.Dup2(int(logFile.Fd()), 2)

	loadDark()
	transforSyntax()
	Display(&buffer)
}
