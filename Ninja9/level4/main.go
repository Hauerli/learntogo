package main

import (
	"fmt"
	"sync"
)

func main() {

	incr := 0
	cnt := 100
	var wg sync.WaitGroup
	var mux sync.Mutex

	wg.Add(cnt)

	for i := 0; i < cnt; i++ {

		go func() {
			mux.Lock()
			v := incr
			v++
			incr = v

			fmt.Println(incr)
			mux.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("final value: ", incr)

}
