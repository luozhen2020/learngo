package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func eval_in_func(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		q, _ := div(a, b)
		return q, nil
	default:
		return 0, fmt.Errorf(
			"Unsupported operation : %s", op)
	}
}

func div(a, b int) (q, r int) {
	q = a / b
	r = a % b
	return
}

/*函数式编程*/
func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf(
		"Calling function %s with args"+
			"(%d, %d)\n", opName, a, b)
	return op(a, b)
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

/*可变参数列表*/
func sumArgs(values ...int) int {
	sum := 0
	for i := range values {
		sum += values[i]
	}
	return sum
}

func swap(a, b *int) {
	*b, *a = *a, *b
}

func swap2(a, b int) (int, int) {
	return b, a
}

func main() {
	if result, err := eval_in_func(3, 4, "x"); err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println(result)
	}
	fmt.Println(div(18, 3))

	fmt.Println(apply(pow, 3, 4))
	/*匿名函数*/
	fmt.Println(apply(func(a, b int) int {
		return int(math.Pow(float64(a), float64(b)))
	}, 3, 4))

	fmt.Println(sumArgs(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))

	a, b := 3, 4
	swap(&a, &b)
	fmt.Println(a, b)

	c, d := swap2(3, 4)
	fmt.Println(c, d)
}
