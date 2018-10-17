package editor

import (
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func isDir(name string) bool {
	fi, err := os.Stat(name)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return fi.IsDir()
}

func (e *Edit) newFileselector(rootDir string) *tview.TreeView {
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
				e.focusMain()

			}
		} else {
			// Collapse if visible, expand if collapsed.
			node.SetExpanded(!node.IsExpanded())
		}
	})
	return tree
}
