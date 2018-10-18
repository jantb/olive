package editor

import (
	"github.com/gdamore/tcell"
	"github.com/jantb/olive/rpc"
	"github.com/jantb/olive/xi"
	"github.com/rivo/tview"
	"io"
	"log"
	"os"
	"time"
)

type Edit struct {
	grid         *tview.Grid
	fileSelector *tview.TreeView
	main         *MainView
	headerTree   *tview.TextView
	header       *tview.TextView
	footerTree   *FooterTree
	footer       *Footer

	Views       map[string]*View
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

func NewEdit(rw io.ReadWriter, configPath string) *Edit {
	e := &Edit{}
	tcell.SetEncodingFallback(tcell.EncodingFallbackASCII)
	e.redraws = make(chan struct{}, 50)
	e.updates = make(chan func(), 1)

	e.Views = make(map[string]*View)
	e.xi = rpc.NewConnection(rw)
	e.xi.Start(configPath)
	e.xi.SetTheme("Solarized (dark)")

	go e.handleRequests()
	go e.mainLoop()
	return e
}

func (e *Edit) mainLoop() {
	for {
		select {
		case update := <-e.updates:
			update()
		case <-e.redraws:
			e.application.Draw()
		}
	}
}
func (e *Edit) focusMain() {
	e.application.SetFocus(e.main)
}

func (e *Edit) focusFileselector() {
	e.application.SetFocus(e.fileSelector)
}

func (e *Edit) OpenFile(path string) {
	msg, err := e.xi.Request(&rpc.Request{
		Method: "new_view",
		Params: &rpc.Object{"file_path": path},
	})
	e.header.SetText(path)
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

func (e *Edit) Init() {

	path := "."
	if len(os.Args) > 1 {
		path = os.Args[1]
	}

	e.fileSelector = e.newFileselector(path)

	e.main = NewMainView()
	e.main.Edit = e
	e.main.dataView = map[string]*Dataview{}
	e.headerTree = tview.NewTextView().
		SetTextAlign(tview.AlignLeft)
	e.header = tview.NewTextView().
		SetTextAlign(tview.AlignLeft)
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

func (e *Edit) Start() {
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

func (e *Edit) handleRequests() {
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
				dataView.ApplyUpdate(msg.Value.(*rpc.Update))
				e.footer.totalLines = dataView.Length()
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
		}

		e.updates <- func() {
			e.redraws <- struct{}{}
		}
	}
}
