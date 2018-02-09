package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 10000000; i ++ {
			fmt.Println(i)
		}
	}()

	//wg.Add(1)
	//go func() {
	//	defer wg.Done()
	//	for i := 0; i < 1000000; i ++ {
	//		fmt.Println(i)
	//	}
	//}()

	wg.Wait()
}
