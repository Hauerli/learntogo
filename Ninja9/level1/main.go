package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func write(text string) {
	fmt.Println(text)
	wg.Done()
}

func main() {

	wg.Add(4)
	for i := 0; i <= 10; i++ {
		go write("test")
	}
	wg.Wait()
}
