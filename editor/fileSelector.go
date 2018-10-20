package editor

import (
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
	"io/ioutil"
	"os"
	"path/filepath"
)

func isDir(name string) bool {
	if fi, err := os.Stat(name); !os.IsNotExist(err) {
		return fi.IsDir()
	}
	if _, err := os.Stat(name); os.IsNotExist(err) {
		os.OpenFile(name, os.O_RDONLY|os.O_CREATE, 0666)
	}
	return false
}

func (e *Editor) newFileselector(rootDir string) *tview.TreeView {
	if !isDir(rootDir) {
		rootDir = "."
	}
	root := tview.NewTreeNode(rootDir).
		SetColor(tcell.ColorRed)
	tree := tview.NewTreeView().
		SetRoot(root).
		SetCurrentNode(root)

	add := func(target *tview.TreeNode, path string) {
		files, err := ioutil.ReadDir(path)
		if err != nil {
			panic(err)
		}
		for _, file := range files {
			node := tview.NewTreeNode(file.Name()).
				SetReference(filepath.Join(path, file.Name())).
				SetSelectable(true)
			if file.IsDir() {
				node.SetColor(tcell.ColorGreen)
			}
			target.AddChild(node)
		}
	}

	add(root, rootDir)

	tree.SetSelectedFunc(func(node *tview.TreeNode) {
		reference := node.GetReference()
		if reference == nil {
			return // Selecting the root node does nothing.
		}
		children := node.GetChildren()
		if len(children) == 0 {
			// Load and show files in this directory.
			path := reference.(string)
			if isDir(path) {
				add(node, path)
			} else {
				e.OpenFile(path)
				e.focusView()
			}
		} else {
			// Collapse if visible, expand if collapsed.
			node.SetExpanded(!node.IsExpanded())
		}
	})
	tree.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch key := event.Key(); key {
		case tcell.KeyEsc:
			if e.view.curViewID != "" {
				e.focusView()
			}
			return nil
		case tcell.KeyRight:
			if e.fileSelector_width < 200 {
				e.fileSelector_width++
			}
			e.application.Draw()
			return nil
		case tcell.KeyLeft:
			if e.fileSelector_width > 3 {
				e.fileSelector_width--
			}
			e.application.Draw()
			return nil
		}
		return event
	})
	return tree
}
