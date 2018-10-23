package editor

import (
	"github.com/atotto/clipboard"
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
	mutex              *sync.Mutex
	grid               *tview.Grid
	fileSelector       *tview.TreeView
	view               *View
	headerTree         *tview.TextView
	header             *Header
	footerTree         *FooterTree
	footer             *Footer
	linenums           *Linenums
	pages              *tview.Pages
	gotoLine           *GotoLine
	openFile           *OpenFile
	linenums_width     int
	fileSelector_width int

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
	e.redraws = make(chan struct{}, 100)
	e.updates = make(chan func(), 1)

	e.xi = rpc.NewConnection(rw)
	e.xi.Start(configPath)
	e.xi.SetTheme("Solarized (dark)")

	go e.handleRequests()
	go e.mainLoop()
	return e
}
func (e *Editor) Quit() {
	dataviews := e.view.dataView
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
			go e.application.Draw()
		}
	}
}
func (e *Editor) focusView() {
	e.application.SetFocus(e.view)
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
	e.view.dataView[viewId] = &Dataview{ViewId: viewId}
	e.view.dataView[viewId].InputHandler = &rpc.InputHandler{ViewID: viewId, FilePath: path, Xi: e.xi}
	e.view.dataView[viewId].LineCache = xi.NewLineCache()
	e.curViewID = viewId
	_, _, w, h := e.view.Box.GetInnerRect()
	e.view.dataView[viewId].Scroll(0, h)
	e.view.dataView[viewId].Resize(w, h)
}

func (e *Editor) updateColumnWidths() {
	e.grid.SetColumns(e.fileSelector_width, e.linenums_width, 0)
}
func (e *Editor) Init() {

	path := "."
	if len(os.Args) > 1 {
		path = os.Args[1]
	}

	e.fileSelector = e.newFileselector(path)
	e.fileSelector_width = 30
	e.view = NewView()
	e.view.Editor = e
	e.view.dataView = map[string]*Dataview{}
	e.headerTree = tview.NewTextView().
		SetTextAlign(tview.AlignLeft)
	e.header = e.NewHeader()
	e.footerTree = e.NewFooterTree()
	e.footer = e.NewFooter()
	e.linenums = e.NewLinenums()
	e.pages = tview.NewPages()

	grid := tview.NewGrid().
		SetRows(1, 0, 1).
		SetColumns(e.fileSelector_width, e.linenums_width, 0).
		SetBorders(false).
		AddItem(e.headerTree, 0, 0, 1, 1, 0, 0, false).
		AddItem(e.header, 0, 1, 1, 2, 0, 0, false).
		AddItem(e.footerTree, 2, 0, 1, 1, 0, 0, false).
		AddItem(e.footer, 2, 1, 1, 2, 0, 0, false)

	grid.AddItem(e.fileSelector, 1, 0, 1, 1, 0, 0, true).
		AddItem(e.linenums, 1, 1, 1, 1, 0, 0, false).
		AddItem(e.view, 1, 2, 1, 1, 0, 0, false)

	e.grid = grid
	modal := getModal()
	e.gotoLine = e.NewGotoLine()
	e.openFile = e.NewOpenFile()

	e.pages.
		AddPage("gotoLine", modal(e.gotoLine, 20, 3), true, false).
		AddPage("openFile", modal(e.openFile, 40, 5), true, false).
		AddPage("editor", e.grid, true, true)
}

func getModal() func(p tview.Primitive, width int, height int) tview.Primitive {
	modal := func(p tview.Primitive, width, height int) tview.Primitive {
		return tview.NewGrid().
			SetColumns(0, width, 0).
			SetRows(0, height, 0).
			AddItem(p, 1, 1, 1, 1, 0, 0, true)
	}
	return modal
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
			e.focusView()
		}()
	}
	e.application.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch key := event.Key(); key {
		case tcell.KeyCtrlC:
			if e.view.HasFocus() {
				dataview := e.view.dataView[e.curViewID]
				data := dataview.Copy()
				if data == "" {
					dataview.MoveToBeginningOfLine()
					dataview.MoveToEndOfLineAndModifySelection()
					data = dataview.Copy()
				}
				clipboard.WriteAll(data)
				return nil
			} else {
				return event
			}
		case tcell.KeyCtrlQ:
			e.application.Stop()
			return nil
		}
		return event
	})

	if err := e.application.SetRoot(e.pages, true).Run(); err != nil {
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
				dataView := e.view.dataView[update.ViewID]
				for dataView == nil {
					dataView = e.view.dataView[update.ViewID]
					time.Sleep(100 * time.Millisecond)
				}
				e.mutex.Lock()
				dataView.ApplyUpdate(msg.Value.(*rpc.Update))
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
				e.view.MakeVisible(scrollTo.Col, scrollTo.Line)
				//Index is 1 based not 0
				e.footer.cursorX = scrollTo.Col + 1
				e.footer.cursorY = scrollTo.Line + 1
			}
		case *rpc.LanguageChanged:
			e.updates <- func() {
				languageChanged := msg.Value.(*rpc.LanguageChanged)
				e.footer.language = languageChanged.LanguageID
			}
		case *rpc.FindStatus:
			e.updates <- func() {
				findStatus := msg.Value.(*rpc.FindStatus)
				e.view.findstatus = findStatus
			}
		}

		e.updates <- func() {
			e.redraws <- struct{}{}
		}
	}
}
