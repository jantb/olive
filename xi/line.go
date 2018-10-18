package xi

type Line struct {
	Text     string
	Cursors  []int
	Styles   []int
	StyleIds map[int][]int
}

func NewLine(text string, cursors []int, styles []int) *Line {
	line := &Line{}
	line.Text = text
	line.Cursors = cursors
	line.Styles = make([]int, 0, 10)
	line.StyleIds = make(map[int][]int)
	line.SetStyles(styles)
	return line
}

func (l *Line) SetStyles(styles []int) {
	offset := 0
	for i := 0; i < len(styles); i += 3 {
		start := offset + styles[i]
		end := start + styles[i+1]
		styleId := styles[i+2]
		for j := start; j <= end; j++ {
			if l.StyleIds[j] == nil {
				l.StyleIds[j] = []int{}
			}
			l.StyleIds[j] = append(l.StyleIds[j], styleId)
		}

		offset = end
	}
}
