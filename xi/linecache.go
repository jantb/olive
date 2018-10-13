package xi

import (
	"github.com/jantb/olive/rpc"
)

type LineCache struct {
	lines         []*Line
	invalidBefore int
	invalidAfter  int
}

func NewLineCache() *LineCache {
	return &LineCache{
		lines: make([]*Line, 0, 1000),
	}
}

func (lc LineCache) InvalidBefore() int {
	return lc.invalidBefore
}

func (lc LineCache) TotalLength() int {
	return lc.invalidBefore + len(lc.lines) + lc.invalidAfter
}

func (lc LineCache) Length() int {
	return lc.invalidBefore + len(lc.lines) + lc.invalidAfter
}

func (lc LineCache) Lines() []*Line {
	return lc.lines
}

func (lc *LineCache) addInvalid(newLines []*Line, newInvalidBefore *int, newInvalidAfter *int, n int) {
	if len(newLines) == 0 {
		*newInvalidBefore += n
	} else {
		*newInvalidAfter += n
	}
}

func (lc *LineCache) ApplyUpdate(update *rpc.Update) {
	newLines := make([]*Line, 0, 1000)
	newInvalidBefore := 0
	newInvalidAfter := 0
	oldIx := 0
	for _, op := range update.Update.Ops {
		switch op.Op {
		case "copy":
			nRemaining := op.N
			if oldIx < lc.invalidBefore {
				nInvalid := 0
				if op.N < lc.invalidBefore-oldIx {
					nInvalid = op.N
				} else {
					nInvalid = lc.invalidBefore - oldIx
				}
				if len(newLines) == 0 {
					newInvalidBefore += nInvalid
				} else {
					newInvalidAfter += nInvalid
				}
				oldIx += nInvalid
				nRemaining -= nInvalid
			}
			if nRemaining > 0 && oldIx < lc.invalidBefore+len(lc.lines) {
				for i := 0; i < newInvalidAfter; i++ {
					newLines = append(newLines, nil)
				}
				newInvalidAfter = 0

				nCopy := 0
				if nRemaining < lc.invalidBefore+len(lc.lines)-oldIx {
					nCopy = nRemaining
				} else {
					nCopy = lc.invalidBefore + len(lc.lines) - oldIx
				}
				start := oldIx - lc.invalidBefore

				newLines = append(newLines, lc.lines[start:start+nCopy]...)
				oldIx += nCopy
				nRemaining -= nCopy
			}
			if len(newLines) == 0 {
				newInvalidBefore += nRemaining
			} else {
				newInvalidAfter += nRemaining
			}
			oldIx += nRemaining

		case "skip":
			oldIx += op.N
		case "invalidate":
			lc.addInvalid(newLines, &newInvalidBefore, &newInvalidAfter, op.N)
		case "ins":
			for i := 0; i < newInvalidAfter; i++ {
				newLines = append(newLines, nil)
			}
			for _, line := range op.Lines {
				newline := NewLine(line.Text, line.Cursor, line.Styles)
				newLines = append(newLines, newline)
			}

		}
	}

	lc.lines = newLines
	lc.invalidBefore = newInvalidBefore
	lc.invalidAfter = newInvalidAfter
}
