package cursor

// Directions to be used with Cursor
const (
	Up = iota
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

func (c *Cursor) canMove(dir int) bool {
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
func (c *Cursor) Move(dir int) {
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
}
