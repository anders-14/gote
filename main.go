package main

import (
	"fmt"
	"log"
	"os"

	"github.com/anders-14/gote/buffer"
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

func handleKeypress(t *term.Term) error {
	c, err := readChar(t)
	if err != nil {
		return err
	}

	switch c {
	case ctrl('q'):
		return fmt.Errorf("ctrl+q pressed, program exits")
	default:
		fmt.Printf("%c\r\n", c)
		break
	}

	return nil
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

	editor.OpenFile("./test.txt")
	fmt.Print(editor.ToString())

	if err := editor.SaveFile(); err != nil {
		log.Printf("[ERR]: %+v\r\n", err)
		return
	}

	for true {
		if err := handleKeypress(t); err != nil {
			log.Printf("[ERR]: %+v\r\n", err)
			return
		}
	}
}
