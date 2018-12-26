package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	incr := 0
	cnt := 100
	var wg sync.WaitGroup

	wg.Add(cnt)

	for i := 0; i < cnt; i++ {

		go func() {
			v := incr
			v++
			runtime.Gosched()
			incr = v
			fmt.Println(incr)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("final value: ", incr)

}
