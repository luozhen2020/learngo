package main

import (
	"fmt"
)

func chanDemo() {
	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i)
	}

	for i := 0; i < 10; i++ {
		workers[i].in <- 'a' + i
		/*<-workers[i].done*/
	}

	for i := 0; i < 10; i++ {
		workers[i].in <- 'A' + i
		/*<-workers[i].done*/
	}

	/*time.Sleep(time.Millisecond)*/

	// wait for all of them
	for _, worker := range workers {
		<-worker.done
		<-worker.done
	}
}

func createWorker(id int) worker {
	w := worker{
		in:   make(chan int),
		done: make(chan bool),
	}
	go doWork(id, w)

	return w
}

type worker struct {
	in   chan int
	done chan bool
}

func doWork(id int, w worker) {
	for n := range w.in {
		/*n, ok := <-c
		if !ok {
			break
		}*/
		fmt.Printf("Worker %d received %c\n", id, n)
		go func() {
			w.done <- true
		}()

	}
}

func main() {
	chanDemo()
}
