package main

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	msg := []byte("hello\n")
	f := bufio.NewWriter(os.Stdout)
	f.Write(msg)
	f.Flush()
	for i := 0; i < 10000000; i ++ {
		fmt.Println("hello")
	}
}
