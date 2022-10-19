package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func tryDefer() {
	/*defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
	panic("error occurred!")
	fmt.Println(4)*/

	for i := 0; i < 100; i++ {
		defer fmt.Println(i)
		if i == 30 {
			panic("printed too many!")
		}
	}
}

func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func writeFile(filename string) {
	//file, err := os.Create(filename)
	file, err := os.OpenFile(filename, os.O_EXCL|os.O_CREATE, 0666)

	err = errors.New("this is a customized error for test!")
	if err != nil {
		//panic(err)
		//fmt.Println("file already exists!")
		/*fmt.Println("Error:", err.Error())*/
		if pathError, ok := err.(*os.PathError); !ok {
			panic(err)
		} else {
			fmt.Printf("%s, %s, %s\n", pathError.Op, pathError.Path, pathError.Err)
		}
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	f := fibonacci()
	for i := 0; i < 30; i++ {
		fmt.Fprintln(writer, f())
	}
}

func main() {
	writeFile("fib.txt")
	//tryDefer()
}
