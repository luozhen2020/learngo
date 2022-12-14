package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 1000; i++ {
		go func(i int) { // race condition!
			for {
				/*a[i]++
				runtime.Gosched()*/
				fmt.Printf("Hello from goroutine %d\n", i)
			}
		}(i)
	}
	time.Sleep(time.Minute)
	//fmt.Println(a)
}
