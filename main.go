package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/pkg/term/termios"
	"golang.org/x/sys/unix"
)

var original_termios unix.Termios

func main() {
	enableRawMode()
	reader := bufio.NewReader(os.Stdin)
	defer func() {
		fmt.Println("disabling raw mode and exiting")
		disableRawMode()
	}()
	for {
		c, err := reader.ReadByte()
		if err != nil || c == 'q' {
			fmt.Printf("\n")
			return 
		}
		fmt.Println(c)
	}
}

func enableRawMode() {
	termios.Tcgetattr(os.Stdin.Fd(), &original_termios)
	
	raw := original_termios
	raw.Lflag &= unix.ECHO

	termios.Tcsetattr(os.Stdin.Fd(), unix.TCSAFLUSH, &raw)
}

func disableRawMode() {
	termios.Tcsetattr(os.Stdin.Fd(), termios.TCSAFLUSH, &original_termios)
}