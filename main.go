package main

import (
	"fmt"
	"log"
	"os"

	"github.com/anders-14/gote/buffer"
	"github.com/anders-14/gote/cursor"
	"github.com/pkg/term"
	"golang.org/x/crypto/ssh/terminal"
)

func ctrl(char byte) byte {
	return char & 31
}

func readChar(t *term.Term) (byte, error) {
	var buf = []byte{'0'}

	_, err := t.Read(buf)
	if err != nil {
		return 0, err
	}

	return buf[0], nil
}

func handleKeypress(t *term.Term, buf *buffer.Buffer) error {
	c, err := readChar(t)
	if err != nil {
		return err
	}

	switch c {
	case ctrl('q'):
		return fmt.Errorf("ctrl+q pressed, program exits")
	case 'h':
		buf.Cursor.Move(cursor.Left)
		break
	case 'j':
		buf.Cursor.Move(cursor.Down)
		break
	case 'k':
		buf.Cursor.Move(cursor.Up)
		break
	case 'l':
		buf.Cursor.Move(cursor.Right)
		break
	default:
		fmt.Printf("%c\r\n", c)
		break
	}

	return nil
}

func draw(buf *buffer.Buffer) {
	fmt.Printf("\x1b[2J")
	fmt.Printf("\x1b[1;1H")
	fmt.Printf(buf.ToString())
	fmt.Printf("\x1b[%d;%dH", buf.Cursor.Y+1, buf.Cursor.X+1)
}

func main() {
	t, _ := term.Open("/dev/tty")
	t.SetRaw()
	defer t.Restore()
	w, h, err := terminal.GetSize(int(os.Stdin.Fd()))
	if err != nil {
		log.Printf("[ERR]: %+v\r\n", err)
		return
	}

	editor := buffer.New(0, 0, w, h, true)

	if len(os.Args) >= 2 {
		if err := editor.OpenFile(os.Args[1]); err != nil {
			log.Printf("[ERR]: %+v\r\n", err)
			return
		}
	}

	for true {
		draw(editor)
		if err := handleKeypress(t, editor); err != nil {
			log.Printf("[ERR]: %+v\r\n", err)
			return
		}
	}
}
