package go_plugin

import (
	"github.com/jantb/olive/ds"
	"log"
	"os/exec"
	"strconv"
	"strings"
)

func Vet() []ds.Position {
	// go get golang.org/x/tools/cmd/goimports
	path, err := exec.LookPath("go")
	if err != nil {
		log.Println("go not installed")
		return []ds.Position{}
	}

	p := exec.Command(path, "vet")

	bytes, _ := p.CombinedOutput()

	s := string(bytes)
	vetErrors := parseVet(s)

	return vetErrors
}

func parseVet(s string) []ds.Position {
	var position []ds.Position
	split := strings.Split(s, "\n")
	for _, vet := range split {
		n := strings.Split(vet, ": ")
		if len(n) != 2 {
			continue
		}
		pathAndPos := n[0]
		pos := strings.Split(pathAndPos, ":")
		column, _ := strconv.Atoi(pos[1])
		row, _ := strconv.Atoi(pos[2])
		position = append(position, ds.Position{Column: column, Row: row, Path: pos[0], Explanation: n[1]})
	}
	return position
}
