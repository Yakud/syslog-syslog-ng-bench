package main

import (
	"bufio"
	"os"
)

func main() {
	msg := []byte("hello\n")
	f := bufio.NewWriter(os.Stdout)
	for i := 0; i < 10000000; i ++ {
		f.Write(msg)
		f.Flush()
	}
}


package main
import (
"os"
)
func main() {
	// create a buffer that has space equal to the page size
	buf := make([]byte, 4096)
	y := []byte("y\n")
	used := 0
	for {
		buf[used] = y[0]
		buf[(used)+1] = y[1]
		used += 2
		if used == 4096 {
			break
		}
	}
	// Print out the entire buffer at once - writing direct to
	// stdout is the fastest way of doing this.
	// It's the size of a standard linux pagefile.
	for {
		os.Stdout.Write(buf)
	}
}