package channel

import (
	"fmt"
	"time"
)

/*(1)Channel 的select运用：select是一种go可以处理多个通道之间的机制，通道(channel)实现了多个goroutine之前的同步或者通信，
  而select则实现了多个通道(channel)的同步或者通信，并且select具有阻塞的特性。*/

// Sample1 - 竞争选举
/*
  最常见的使用场景，多个通道，有一个满足条件可以读取，就可以“竞选成功”。
*/
func Sample1() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	go func() {
		time.Sleep(time.Second)
		ch1 <- 1
	}()

	go func() {
		ch2 <- 2
	}()

	go func() {
		ch3 <- 3
	}()

	select {
	case i := <-ch1:
		fmt.Printf("从ch1读取了数据%d", i)
	case j := <-ch2:
		fmt.Printf("从ch2读取了数据%d", j)
	case m := <-ch3:
		fmt.Printf("从ch3读取了数据%d", m)
		//...
	}
}

// Sample2 - 超时处理（保证不阻塞）
/*
  因为select是阻塞的，我们有时候就需要搭配超时处理来处理这种情况，超过某一个时间就要进行处理，保证程序不阻塞。
*/
func Sample2() {
	ch1 := make(chan byte)

	select {
	case str := <-ch1:
		fmt.Printf("received str: %s", string(str))
	case <-time.After(time.Second * 5):
		fmt.Println("timeout!!")
	}
}

// Sample3 - 判断buffered channel是否阻塞
/*
  这个例子很经典，比如我们有一个有限的资源（这里用buffer channel实现），我们每一秒向bufChan传送数据，
  由于生产者的生产速度大于消费者的消费速度，故会触发default语句，这个就很像我们web端来显示并发过高的提示了，
  小伙伴们可以尝试删除go func中的time.Sleep(5*time.Second)，看看是否还会触发default语句?
*/
func Sample3() {
	bufChan := make(chan int, 5)

	go func() {
		time.Sleep(time.Second)
		for {
			<-bufChan
			time.Sleep(5 * time.Second)
		}
	}()

	iCount := 1
	for {
		select {
		case bufChan <- iCount:
			fmt.Printf("add iCount #%d success\n", iCount)
			time.Sleep(time.Second)
		default:
			fmt.Println("资源已满，请稍后再试")
			time.Sleep(time.Second)
		}
		iCount++
	}
}

// Sample4 - 阻塞main函数
/*
  有时候我们会让main函数阻塞不退出，如http服务，我们会使用空的select{}来阻塞main goroutine
*/
func Sample4() {
	ch := make(chan int)
	go func() {
		for {
			ch <- 1 // 生产
			time.Sleep(time.Second)
		}
	}()
	go func() {
		for {
			fmt.Println(<-ch) // 消费
		}
	}()
	select {} // 无条件死锁，否则程序直接退出看不到打印结果
}
