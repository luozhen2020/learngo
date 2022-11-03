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
	for n := range c {
		/*n, ok := <-c
		if !ok {
			break
		}*/
		fmt.Printf("Worker %d received %c\n", id, n)
	}
}

func bufferedChannel() {
	c := make(chan int, 3)

	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	c <- 'e'

	time.Sleep(time.Millisecond)
}

func channelClose() {
	c := make(chan int)

	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	c <- 'e'

	close(c)
	time.Sleep(time.Millisecond)
}

func main() {
	fmt.Println("Channel as first-class citizen")
	chanDemo()
	fmt.Println("Buffered channel")
	bufferedChannel()
	fmt.Println("Channel close and range")
	channelClose()
}
