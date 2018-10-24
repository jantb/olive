package editor

import (
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type OpenFile struct {
	*tview.Box
	*Editor
	input  string
	result []string
	index  int
	files  []File
}

// NewView returns a new view view primitive.
func (e *Editor) NewOpenFile() *OpenFile {
	file := OpenFile{
		Box:    tview.NewBox().SetBorder(false),
		Editor: e,
		files:  make([]File, 10),
		result: make([]string, 10),
	}
	file.traversePath()
	return &file

}

type File struct {
	FilePath string
}

func (g *OpenFile) traversePath() {
	err := filepath.Walk(".", g.visit)
	if err != nil {
		log.Fatal(err)
	}
}

func (g *OpenFile) visit(path string, f os.FileInfo, err error) error {
	log.Println(f.Name())
	if f.Name() == ".git" || f.Name() == ".idea" {
		return filepath.SkipDir
	}
	if !f.IsDir() {
		file := File{FilePath: path}
		g.files = append(g.files, file)
	}
	return nil
}

// Draw draws this primitive onto the screen.
func (g *OpenFile) Draw(screen tcell.Screen) {
	_, bg, _ := defaultStyle.Decompose()
	g.Box.SetBackgroundColor(bg).SetTitle(" Open file ").SetTitleAlign(tview.AlignLeft).SetBorder(true).Draw(screen)
	offx := g.drawText(screen, "Enter file name: ", 0, 0, defaultStyle.Foreground(tcell.ColorLightCyan))
	offx = g.drawText(screen, g.input, offx, 0, defaultStyle.Background(tcell.ColorDarkCyan).Foreground(tcell.ColorYellow))
	for key, value := range g.result[:Min(len(g.result), 10)] {
		if key == g.index {
			g.drawText(screen, value, 0, key+1, defaultStyle.Background(tcell.ColorLightCyan).Foreground(tcell.ColorBlack))
		} else {
			g.drawText(screen, value, 0, key+1, defaultStyle.Background(tcell.ColorDarkCyan).Foreground(tcell.ColorYellow))
		}
	}
}

func (g *OpenFile) drawText(screen tcell.Screen, text string, offsetX, offsetY int, style tcell.Style) (offset int) {
	for x, r := range text {
		offset = g.draw(screen, x+offsetX, offsetY, r, style)
	}
	return offset
}

func (g *OpenFile) draw(screen tcell.Screen, x, y int, r rune, style tcell.Style) int {
	xr, yr, _, _ := g.Box.GetInnerRect()
	screen.SetContent(xr+x, yr+y, r, nil, style)
	return x + 1
}

func (g *OpenFile) InputHandler() func(event *tcell.EventKey, setFocus func(p tview.Primitive)) {

	return g.WrapInputHandler(func(key *tcell.EventKey, setFocus func(p tview.Primitive)) {
		switch key.Key() {
		case tcell.KeyEsc:
			g.pages.HidePage("openFile")
			g.focusView()
		case tcell.KeyRune:
			r := key.Rune()
			if len(g.input) < 20 {
				g.input = g.input + string(r)
				g.result = g.result[:0]
				g.index = 0
				for _, value := range g.files {
					if strings.Contains(value.FilePath, g.input) {
						g.result = append(g.result, value.FilePath)
					}
				}
			}
		case tcell.KeyBackspace2:
			if len(g.input) > 0 {
				g.input = g.input[:len(g.input)-1]
				g.result = g.result[:0]
				g.index = 0
				for _, value := range g.files {
					if strings.Contains(value.FilePath, g.input) {
						g.result = append(g.result, value.FilePath)
					}
				}
			}
		case tcell.KeyDown:
			if g.index+1 < len(g.result[:Min(len(g.result), 10)]) {
				g.index++
			}

		case tcell.KeyUp:
			if g.index > 0 && g.index <= len(g.result[:Min(len(g.result), 10)]) {
				g.index--
			}
		case tcell.KeyEnter:
			if len(g.result) == 0 {
				return
			}
			g.pages.HidePage("openFile")
			g.focusView()
			g.view.OpenFile(g.result[g.index])
			g.input = ""
		}
	})
}
