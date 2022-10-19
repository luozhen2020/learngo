package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func div(a, b int) (int, int) {
	return a / b, a % b
}

func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args "+
		"(%d %d)\n", opName, a, b)
	return op(a, b)
}

//Override math.pow
func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func main() {
	fmt.Println(div(15, 9))

	fmt.Println(apply(pow, 3, 4))
}
