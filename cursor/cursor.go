package cursor

import "fmt"

type direction int

// Directions to be used with Cursor
const (
	Up direction = iota
	Down
	Left
	Right
)

// Cursor keeps track of the cursor position and bounds
type Cursor struct {
	X, Y       int
	xmax, ymax int
}

// New creates a new Cursor given Buffer size in rows and col
func New(rows, cols int) *Cursor {
	return &Cursor{
		X:    0,
		Y:    0,
		xmax: cols - 1,
		ymax: rows - 1,
	}
}

func (c *Cursor) canMove(dir direction) bool {
	switch dir {
	case Up:
		return c.Y > 0
	case Down:
		return c.Y < c.ymax
	case Left:
		return c.X > 0
	case Right:
		return c.X < c.xmax
	}

	return false
}

// Move moves the cursor a given direction
func (c *Cursor) Move(dir direction) error {
	if dir < 0 || dir > 3 {
		return fmt.Errorf("invalid direction: %d", dir)
	}

	if c.canMove(dir) {
		switch dir {
		case Up:
			c.Y--
			break
		case Down:
			c.Y++
			break
		case Left:
			c.X--
			break
		case Right:
			c.X++
			break
		}
	}

	return nil
}
