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
