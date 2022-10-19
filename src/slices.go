package main

import "fmt"

func updateSlice(s []int) {
	s[0] = 100
}

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}

	fmt.Println("arr[2:6]=", arr[2:6])
	fmt.Println("arr[2:]=", arr[2:])
	fmt.Println("arr[:6]=", arr[:6])
	fmt.Println("arr[:]=", arr[:])

	s1 := arr[2:]
	s2 := arr[:]
	updateSlice(s1)
	updateSlice(s2)
	fmt.Println("Updating s1=", s1)
	fmt.Println("Updating s2=", s2)

	fmt.Println("Reslice operation>>>")
	s2 = s2[:5]
	fmt.Println("After reslice1, s2 =", s2)
	s2 = s2[2:]
	fmt.Println("After reslice2, s2 =", s2)

	fmt.Println("Extending slice>>>")
	arr[0], arr[2] = 0, 2
	s1 = arr[2:6]
	s2 = s1[3:5] //s1[3], s1[4]
	//fmt.Println("s1[4]=", s1[4])
	fmt.Println("s1=", s1)
	fmt.Println("s2=", s2)
}
