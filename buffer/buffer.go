package buffer

import (
	"github.com/anders-14/gote/cursor"
)

// Buffer hold text to be displayed on screen
// also keeps track of a cursor, the buffers
// position on screen and if the buffer currently
// is in focus
type Buffer struct {
	Cursor        *cursor.Cursor
	Rows          [][]byte
	x, y          int
	width, height int
	focus         bool
}

// New creates a new Buffer given its position, size and
// if it should be focused by default
func New(x, y, width, height int, focus bool) *Buffer {
	return &Buffer{
		Cursor: cursor.New(height, width),
		Rows:   [][]byte{},
		x:      x,
		y:      y,
		width:  width,
		height: height,
		focus:  focus,
	}
}

// ToString converts Buffer.Rows to a string
func (b *Buffer) ToString() string {
	buf := ""

	for _, row := range b.Rows {
		buf += string(row) + "\r\n"
	}

	return buf
}

// AppendRow add a new row to the end of the Buffer
// containing the given chars
func (b *Buffer) AppendRow(chars []byte) {
	b.Rows = append(b.Rows, chars)
}

// Insert inserts char at Buffer.rows[rowIdx][colIdx]
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
