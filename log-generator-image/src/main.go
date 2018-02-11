package main

import (
	//"os"
	//"bufio"
	//"fmt"
	"os"
)

func main() {
	//f := bufio.NewWriter(os.Stdout)
	//for i := 0; i < 10000000; i ++ {
	//	fmt.Println("yes")
	//	//f.WriteString("yes\n")
	//	//f.Flush()
	//}

	// create a buffer that has space equal to the page size
	//buf := make([]byte, 4096)
	//y := []byte("y\n")
	//used := 0
	//for {
	//	buf[used] = y[0]
	//	buf[(used)+1] = y[1]
	//	used += 2
	//	if used == 4096 {
	//		break
	//	}
	//}
	// Print out the entire buffer at once - writing direct to
	// stdout is the fastest way of doing this.
	// It's the size of a standard linux pagefile.
	msg := []byte("hello\n")
	for {
		os.Stdout.Write(msg)
	}
}
