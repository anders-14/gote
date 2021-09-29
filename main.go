package main

import (
	"fmt"
	"log"

	"github.com/pkg/term"
)

func ctrl(char byte) byte {
	return char & 31
}

func readChar(t *term.Term) (byte, error) {
	var buf = []byte{'0'}
	c := &buf[0]

	_, err := t.Read(buf)
	if err != nil {
		return 0, err
	}

	switch *c {
	case ctrl('q'):
		return 0, fmt.Errorf("program exit")
	}

	return *c, nil
}

func handleKeypress(t *term.Term) error {
	c, err := readChar(t)
	if err != nil {
		return err
	}

	fmt.Printf("%c\r\n", c)

	return nil
}

func main() {
	t, _ := term.Open("/dev/tty")
	t.SetRaw()
	defer t.Restore()

	for true {
		err := handleKeypress(t)
		if err != nil {
			log.Printf("[ERR]: %+v\r\n", err)
			break
		}
	}

}
