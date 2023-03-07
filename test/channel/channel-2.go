package channel

import (
	"fmt"
	"time"
)

/*(2)Channel 的应用场景：
  以下几种情况，当前协程直接阻塞：
    读写无缓存通道;
    写已满的有缓存通道；
    读已空的有缓存通道；

  阻塞时，当前协程暂停执行，直到另一个协程通过写入或读取来解除阻塞；
  如果没有协程来解除阻塞，也就是只有消费或只有生产，就造成所谓的死锁；*/

// Sample5 - 使用 for range 读通道
/*
  当需要不断从通道读取数据时，用 for range 读取通道既安全又便利；
  当通道关闭时，循环会自动退出；
  但也要注意，若 range 没有被正确退出，会导致死锁错误；
*/
func Sample5() {
	ch := make(chan int, 3)

	ch <- 0
	ch <- 1
	ch <- 2

	for x := range ch {
		fmt.Println(x)
		if len(ch) == 0 {
			close(ch)
		}
	}
}

// Sample6 - 使用 select 处理多个通道的数据
/*
  当需要用一个过程处理来自多个通道的数据时可以用 select；
  select 能同时监听多个通道；
  每个 case 都必须是一个通道表达式；
  所有通道表达式都会被求值；
  若只有一个 case 激活（不阻塞），则执行这个 case 块；
  若有多个 case 激活，则随机挑选一个 case 执行；
  若所有 case 都阻塞，且定义了 default 模块，则执行 default 模块；
  若所有 case 都阻塞，且未定义 default 模块，则 select 阻塞，直到有 case 激活；
  一个 case 执行完成后会跳出 select 块；
  当通道为 nil 时，对应的 case 永远为阻塞，相当于删除了该 case；
  用 break label 跳出 for select 结构；
*/
func Sample6() {
	ch1 := make(chan int, 3)
	ch2 := make(chan int, 3)

	go func() {
		for {
			select {
			case x := <-ch1:
				fmt.Println("ch1: ", x)
			case y := <-ch2:
				fmt.Println("ch2: ", y)
			case <-time.After(time.Second * 5):
				fmt.Println("timeout!!")
				return
			}
		}
	}()

	ch1 <- 1
	ch1 <- 2
	ch1 <- 3

	ch2 <- 1
	ch2 <- 2
	ch2 <- 3

	select {
	case <-time.After(time.Second * 5):
		close(ch1)
		close(ch2)
	}
}

// Sample7 - 空闲超时
/*
  基于 select 的特性，可以方便的实现空闲超时；
  方法是在一个 case 上放 time.After()，其执行条件为：若该 case 先激活及超时；
*/
func Sample7() {
	ch1 := make(chan int, 3)
	ch2 := make(chan int, 3)

	go func() {
		for {
			select {
			case x := <-ch1:
				fmt.Println("ch1:", x)
			case y := <-ch2:
				fmt.Println("ch2:", y)
			case <-time.After(time.Second * 1): // 全部阻塞 1 秒后会触发
				fmt.Println("timeout")
			}
		}
	}()

	select {}
}

// Sample8 - 空闲超时
/*
  所有读一个通道的协程，都会收到这个通道的 close() 信号；
  利用这个广播特性可以方便地实现对所有相关协程的集体关闭；
*/
func Sample8() {
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)
	chQuit := make(chan int)

	go func() {
		for {
			select {
			case x := <-ch1:
				fmt.Println("ch1:", x)
			case <-chQuit:
				fmt.Println("ch1 quit")
				return
			}
		}
	}()

	go func() {
		for {
			select {
			case y := <-ch2:
				fmt.Println("ch2:", y)
			case <-chQuit:
				fmt.Println("ch2 quit")
				return
			}
		}
	}()

	time.Sleep(time.Millisecond * 1000)

	close(chQuit)

	select {}
}
