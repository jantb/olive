package editor

import (
	"github.com/jantb/olive/rpc"
	"github.com/jantb/olive/xi"
)

type Dataview struct {
	*xi.LineCache
	*rpc.InputHandler
	ViewId string
}
