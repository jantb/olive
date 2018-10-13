package editor

import (
	"fmt"
	"github.com/gdamore/tcell"
	"github.com/jantb/olive/rpc"
	"io"
	"log"
	"os"
)

type Editor struct {
	screen    tcell.Screen
	Views     map[string]*View
	curViewID string
	xi        *rpc.Connection

	// ui events
	events chan tcell.Event
	// user events
	redraws chan struct{}
	updates chan func()
}

func (e *Editor) CurView() *View {
	return e.Views[e.curViewID]
}

func (e *Editor) initScreen() {
	var err error

	e.screen, err = tcell.NewScreen()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err = e.screen.Init(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if e.screen.HasMouse() {
		e.screen.EnableMouse()
	}
	e.screen.Clear()
}

func (e *Editor) handleEvent(ev tcell.Event) {
	switch ev.(type) {
	case *tcell.EventKey:
		e.CurView().HandleEvent(ev)
	case *tcell.EventMouse:
		e.CurView().HandleEvent(ev)
	}
}

func NewEditor(rw io.ReadWriter, configPath string) *Editor {
	e := &Editor{}

	tcell.SetEncodingFallback(tcell.EncodingFallbackASCII)

	// screen event channel
	e.events = make(chan tcell.Event, 50)
	e.redraws = make(chan struct{}, 50)
	e.updates = make(chan func(), 1)

	e.Views = make(map[string]*View)

	e.xi = rpc.NewConnection(rw)
	e.xi.Start(configPath)
	e.xi.SetTheme("Solarized (dark)")

	return e
}

func (e *Editor) CloseView(v *View) {
	delete(e.Views, v.ID)
}

func (e *Editor) handleRequests() {
	for {
		msg := <-e.xi.Messages

		switch msg.Value.(type) {
		case *rpc.Update:
			e.updates <- func() {
				update := msg.Value.(*rpc.Update)
				view := e.Views[update.ViewID]
				view.ApplyUpdate(msg.Value.(*rpc.Update))
			}
		case *rpc.DefineStyle:
			e.updates <- func() {
				styles.defineStyle(msg.Value.(*rpc.DefineStyle))
			}
		case *rpc.ThemeChanged:
			e.updates <- func() {
				themeChanged := msg.Value.(*rpc.ThemeChanged)
				theme := themeChanged.Theme

				bg := tcell.NewRGBColor(theme.Bg.ToRGB())
				defaultStyle = defaultStyle.Background(bg)
				fg := tcell.NewRGBColor(theme.Fg.ToRGB())
				defaultStyle = defaultStyle.Foreground(fg)

				e.screen.SetStyle(defaultStyle)

				log.Printf("Theme:%v", theme)
			}
		case *rpc.ScrollTo:
			e.updates <- func() {
				scrollTo := msg.Value.(*rpc.ScrollTo)
				e.Views[scrollTo.ViewID].MakeLineVisible(scrollTo.Line, scrollTo.Col)
				e.Views[scrollTo.ViewID].view.ShowCursor(scrollTo.Col, scrollTo.Line)
			}
		}

		e.updates <- func() {
			e.redraws <- struct{}{}
		}
	}
}

func (e *Editor) Start() {
	// Exit gracefully when no filename is provided.
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "No filename provided. Example: `olive example.txt`\n")
		os.Exit(1)
	}

	e.initScreen()
	defer e.screen.Fini()

	quit := make(chan bool, 1)

	go func() {
		for {
			if e.screen != nil {
				// feed events into channel
				e.events <- e.screen.PollEvent()
			}
		}
	}()

	path := os.Args[1]
	vp := NewViewport(e.screen, 0, 0)
	vp.FillParent()
	view, _ := NewView(path, vp, e.xi)
	e.Views[view.ID] = view
	e.curViewID = view.ID

	go e.handleRequests()

	// editor loop
	for {
		if len(e.Views) != 0 {
			curView := e.CurView()
			e.screen.Clear()
			curView.Draw()
			e.screen.Show()
		} else {
			quit <- true
		}
		var event tcell.Event
		select {
		case event = <-e.events:
		case update := <-e.updates:
			update()
		case <-e.redraws:

		case <-quit:
			e.screen.Fini()
			os.Exit(0)
		}

		for event != nil {
			switch ev := event.(type) {
			case *tcell.EventKey:
				switch ev.Key() {
				case tcell.KeyCtrlQ:
					for _, view := range e.Views {
						view.Close()
					}
					close(quit)
				}
			case *tcell.EventResize:
				e.redraws <- struct{}{}
				e.screen.Sync()
			}

			e.handleEvent(event)

			// continue handling events
			select {
			case event = <-e.events:
			default:
				event = nil
			}
		}
	}
}
