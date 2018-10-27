package go_plugin

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	s := `editor/gutter.go:39:34: l.Editor.view.offy undefined (type *View has no field or method offy)
editor/linenums.go:35:35: l.Editor.view.offy undefined (type *View has no field or method offy)
editor/view.go:57:11: v.offy undefined (type *View has no field or method offy)
editor/view.go:128:30: v.offy undefined (type *View has no field or method offy)
editor/view.go:172:11: v.offy undefined (type *View has no field or method offy)
editor/view.go:173:4: v.offy undefined (type *View has no field or method offy)
editor/view.go:176:20: v.offy undefined (type *View has no field or method offy)
editor/view.go:177:4: v.offy undefined (type *View has no field or method offy)
`
	for _, value := range parseVet(s) {
		fmt.Println(value)
	}

}
