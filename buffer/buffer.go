package buffer

import (
	"github.com/anders-14/gote/cursor"
)

type Buffer struct {
	Cursor        *cursor.Cursor
	Rows          [][]byte
	x, y          int
	width, height int
}

func New(x, y, width, height int) *Buffer {
	return &Buffer{
		Cursor: cursor.New(height, width),
		Rows:   [][]byte{},
		x:      x,
		y:      y,
		width:  width,
		height: height,
	}
}

func (b *Buffer) ToString() string {
	buf := ""

	for _, row := range b.Rows {
		buf += string(row) + "\r\n"
	}

	return buf
}

func (b *Buffer) AppendRow(row []byte) {
	b.Rows = append(b.Rows, row)
}

func (b *Buffer) Insert(char byte, rowIdx, colIdx int) {
	if rowIdx >= len(b.Rows) {
		return
	}
	if colIdx >= len(b.Rows[rowIdx]) {
		return
	}

	b.Rows[rowIdx] = append(b.Rows[rowIdx], 0)
	copy(b.Rows[rowIdx][colIdx+1:], b.Rows[rowIdx][colIdx:])
	b.Rows[rowIdx][colIdx] = char
}
