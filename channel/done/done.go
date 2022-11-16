package main

import (
	"fmt"
	"sync"
)

/*func chanDemo() {
	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i)
	}

	for i := 0; i < 10; i++ {
		workers[i].in <- 'a' + i
		//<-workers[i].done
	}

	for _, worker := range workers {
		<-worker.done
	}

	for i := 0; i < 10; i++ {
		workers[i].in <- 'A' + i
		//<-workers[i].done
	}

	//time.Sleep(time.Millisecond)

	// wait for all of them
	for _, worker := range workers {
		<-worker.done
	}
}*/

func chanDemo() {
	var workers [10]worker
	group := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i, &group)
	}
	group.Add(20)

	for i := 0; i < 10; i++ {
		workers[i].in <- 'a' + i
	}

	for i := 0; i < 10; i++ {
		workers[i].in <- 'A' + i
	}

	group.Wait()
}

/*func createWorker(id int) worker {
	w := worker{
		in:   make(chan int),
		done: make(chan bool),
	}
	go doWork(id, w)

	return w
}*/

func createWorker(id int, wg *sync.WaitGroup) worker {
	w := worker{
		in: make(chan int),
		done: func() {
			wg.Done()
		},
	}
	go doWork(id, w)

	return w
}

/*type worker struct {
	in   chan int
	done chan bool
}*/

type worker struct {
	in   chan int
	done func()
}

/*func doWork(id int, w worker) {
	for n := range w.in {
		fmt.Printf("Worker %d received %c\n", id, n)
		w.done <- true

	}
}*/

func doWork(id int, w worker) {
	for n := range w.in {
		fmt.Printf("Worker %d received %c\n", id, n)
		w.done()
	}
}

func main() {
	chanDemo()
}
