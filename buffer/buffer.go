package buffer

import (
	"os"
	"strings"

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

// OpenFile opens a file given the name and reads it into the buffer
func (b *Buffer) OpenFile(filename string) error {
	contents, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	rows := strings.Split(string(contents), "\n")
	for i, row := range rows {
		if i != len(rows)-1 {
			b.AppendRow([]byte(row))
		}
	}

	return nil
}
