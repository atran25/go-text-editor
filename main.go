package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/pkg/term/termios"
	"golang.org/x/sys/unix"
)

func main() {
	enableRawMode()
	reader := bufio.NewReader(os.Stdin)
	for {
		c, err := reader.ReadByte()
		if err != nil || c == 'q' {
			return
		}
		fmt.Println(c)
	}
}

func enableRawMode() {
	var raw unix.Termios
	termios.Tcgetattr(os.Stdin.Fd(), &raw)
	
	raw.Lflag &= unix.ECHO

	termios.Tcsetattr(os.Stdin.Fd(), unix.TCSAFLUSH, &raw)
}
