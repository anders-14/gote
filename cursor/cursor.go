package cursor

const (
	UP = iota
	DOWN
	LEFT
	RIGHT
)

type Cursor struct {
	X, Y       int
	xmax, ymax int
}

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
	case UP:
		return c.Y > 0
	case DOWN:
		return c.Y < c.ymax
	case LEFT:
		return c.X > 0
	case RIGHT:
		return c.X < c.xmax
	}

	return false
}

func (c *Cursor) Move(dir int) {
	if c.canMove(dir) {
		switch dir {
		case UP:
			c.Y--
			break
		case DOWN:
			c.Y++
			break
		case LEFT:
			c.X--
			break
		case RIGHT:
			c.X++
			break
		}
	}
}
