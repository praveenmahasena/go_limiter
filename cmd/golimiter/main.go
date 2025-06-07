// package main
package main

import (
	"fmt"
	"os"
	"syscall"

	"github.com/praveenmahasena/go-limiter/internal"
)

func main() {
	if err := internal.Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		syscall.Exit(-1)
	}
}
