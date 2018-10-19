package editor

import (
	"github.com/gdamore/tcell"
	"github.com/jantb/olive/rpc"
	"github.com/jantb/olive/xi"
	"github.com/rivo/tview"
	"io"
	"log"
	"os"
	"sync"
	"time"
)

type Editor struct {
	mutex        *sync.Mutex
	grid         *tview.Grid
	fileSelector *tview.TreeView
	main         *View
	headerTree   *tview.TextView
	header       *Header
	footerTree   *FooterTree
	footer       *Footer

	curViewID   string
	xi          *rpc.Connection
	application *tview.Application
	theme       *rpc.Theme
	// ui events
	events chan tcell.Event
	// user events
	redraws chan struct{}
	updates chan func()
}

func NewEdit(rw io.ReadWriter, configPath string) *Editor {
	e := &Editor{}
	e.mutex = &sync.Mutex{}
	tcell.SetEncodingFallback(tcell.EncodingFallbackASCII)
	e.redraws = make(chan struct{}, 1)
	e.updates = make(chan func(), 1)

	e.xi = rpc.NewConnection(rw)
	e.xi.Start(configPath)
	e.xi.SetTheme("Solarized (dark)")

	go e.handleRequests()
	go e.mainLoop()
	return e
}
func (e *Editor) Quit() {
	dataviews := e.main.dataView
	for _, value := range dataviews {
		value.Close()
	}
}

func (e *Editor) mainLoop() {
	for {
		select {
		case update := <-e.updates:
			update()
		case <-e.redraws:
			e.application.Draw()
		}
	}
}
func (e *Editor) focusMain() {
	e.application.SetFocus(e.main)
}

func (e *Editor) focusFileselector() {
	e.application.SetFocus(e.fileSelector)
}

func (e *Editor) OpenFile(path string) {
	msg, err := e.xi.Request(&rpc.Request{
		Method: "new_view",
		Params: &rpc.Object{"file_path": path},
	})
	e.header.path = path
	if err != nil {
		log.Fatal(err)
	}

	viewId := msg.Value.(string)
	e.main.dataView[viewId] = &Dataview{ViewId: viewId}
	e.main.dataView[viewId].InputHandler = &rpc.InputHandler{ViewID: viewId, FilePath: path, Xi: e.xi}
	e.main.dataView[viewId].LineCache = xi.NewLineCache()
	e.curViewID = viewId
	_, _, w, h := e.main.Box.GetInnerRect()
	e.main.dataView[viewId].Scroll(0, h)
	e.main.dataView[viewId].Resize(w, h)
}

func (e *Editor) Init() {

	path := "."
	if len(os.Args) > 1 {
		path = os.Args[1]
	}

	e.fileSelector = e.newFileselector(path)

	e.main = NewView()
	e.main.Editor = e
	e.main.dataView = map[string]*Dataview{}
	e.headerTree = tview.NewTextView().
		SetTextAlign(tview.AlignLeft)
	e.header = e.NewHeader()
	e.footerTree = e.NewFooterTree()
	e.footer = e.NewFooter()
	grid := tview.NewGrid().
		SetRows(1, 0, 1).
		SetColumns(30, 0).
		SetBorders(false).
		AddItem(e.headerTree, 0, 0, 1, 1, 0, 0, false).
		AddItem(e.header, 0, 1, 1, 1, 0, 0, false).
		AddItem(e.footerTree, 2, 0, 1, 1, 0, 0, false).
		AddItem(e.footer, 2, 1, 1, 1, 0, 0, false)

	// Layout for screens narrower than 100 cells (fileSelector and side bar are hidden).
	grid.AddItem(e.fileSelector, 1, 0, 1, 1, 0, 0, true).
		AddItem(e.main, 1, 1, 1, 1, 0, 0, false)

	e.grid = grid
}

func (e *Editor) Start() {
	e.Init()
	e.application = tview.NewApplication()
	path := "."
	if len(os.Args) > 1 {
		path = os.Args[1]
	}
	if !isDir(path) {
		// Todo hack
		go func() {
			time.Sleep(100 * time.Millisecond)
			e.OpenFile(path)
			e.focusMain()
		}()
	}
	if err := e.application.SetRoot(e.grid, true).Run(); err != nil {
		panic(err)
	}
}

func (e *Editor) handleRequests() {
	for {
		msg := <-e.xi.Messages
		switch msg.Value.(type) {
		case *rpc.Update:
			e.updates <- func() {
				update := msg.Value.(*rpc.Update)
				dataView := e.main.dataView[update.ViewID]
				for dataView == nil {
					dataView = e.main.dataView[update.ViewID]
					time.Sleep(100 * time.Millisecond)
				}

				log.Println("trying to get Lock in update")
				e.mutex.Lock()
				log.Println("Got Lock in update")
				dataView.ApplyUpdate(msg.Value.(*rpc.Update))
				log.Println("UnLocking in update")
				e.mutex.Unlock()
				e.footer.totalLines = dataView.Length()
				e.header.pristine = update.Update.Pristine
			}
		case *rpc.DefineStyle:
			e.updates <- func() {
				styles.defineStyle(msg.Value.(*rpc.DefineStyle))
			}
		case *rpc.ThemeChanged:
			e.updates <- func() {
				themeChanged := msg.Value.(*rpc.ThemeChanged)
				e.theme = &themeChanged.Theme

				bg := tcell.NewRGBColor(e.theme.Bg.ToRGB())
				defaultStyle = defaultStyle.Background(bg)
				fg := tcell.NewRGBColor(e.theme.Fg.ToRGB())
				defaultStyle = defaultStyle.Foreground(fg)

				e.headerTree.SetBackgroundColor(tcell.NewRGBColor(e.theme.Bg.ToRGB()))
				e.headerTree.SetBackgroundColor(tcell.NewRGBColor(e.theme.Bg.ToRGB()))
				e.header.SetBackgroundColor(tcell.NewRGBColor(e.theme.Bg.ToRGB()))

				e.application.SetBeforeDrawFunc(func(screen tcell.Screen) bool {
					screen.SetStyle(defaultStyle)
					return false
				})
			}
		case *rpc.ScrollTo:
			e.updates <- func() {
				scrollTo := msg.Value.(*rpc.ScrollTo)
				e.main.MakeVisible(scrollTo.Col, scrollTo.Line)
				e.footer.cursorX = scrollTo.Col
				e.footer.cursorY = scrollTo.Line
			}
		case *rpc.LanguageChanged:
			e.updates <- func() {
				log.Println("lan")
				languageChanged := msg.Value.(*rpc.LanguageChanged)
				e.footer.language = languageChanged.LanguageID
			}
		}

		e.updates <- func() {
			e.redraws <- struct{}{}
		}
	}
}
