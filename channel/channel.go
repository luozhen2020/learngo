package main

import (
	"fmt"
	"time"
)

func chanDemo() {
	var channels [10]chan<- int
	for i := 0; i < 10; i++ {
		channels[i] = createWorker(i)
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}

	time.Sleep(time.Millisecond)
}

func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)

	return c
}

func worker(id int, c chan int) {
	for {
		fmt.Printf("Worker %d received %c\n", id, <-c)
	}
}

func bufferedChannel() {
	c := make(chan int, 3)

	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'

	time.Sleep(time.Millisecond)
}

func main() {
	//chanDemo()
	bufferedChannel()
}
