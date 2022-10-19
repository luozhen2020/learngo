package main

import "fmt"

func updateSlices(s []int) {
	s[0] = 100
}

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println("arr[2:6]=", arr[2:6])
	fmt.Println("arr[:6]=", arr[:6])
	s1 := arr[2:]
	fmt.Println("s1=", s1)
	s2 := arr[:]
	fmt.Println("s2=", s2)

	fmt.Println("After updateSlice(s1)")
	updateSlices(s1)
	fmt.Println(s1)
	fmt.Println(arr)

	fmt.Println("After updateSlice(s2)")
	updateSlices(s2)
	fmt.Println(s2)
	fmt.Println(arr)

	fmt.Println("Extending slice...")
	arr[0], arr[2] = 0, 2
	s1 = arr[2:6]
	s2 = s1[3:5]
	fmt.Printf("s1=%v, len(s1)=%d, cap(s1)=%d\n", s1, len(s1), cap(s1))
	fmt.Printf("s2=%v, len(s2)=%d, cap(s2)=%d\n", s2, len(s2), cap(s2))

	s3 := append(s2, 10)
	s4 := append(s3, 11)
	s5 := append(s4, 12)
	fmt.Printf("s3=%v, len(s3)=%d, cap(s3)=%d\n", s3, len(s3), cap(s3))
	fmt.Printf("s4=%v, len(s4)=%d, cap(s4)=%d\n", s4, len(s4), cap(s4))
	fmt.Printf("s5=%v, len(s5)=%d, cap(s5)=%d\n", s5, len(s5), cap(s5))
	fmt.Println("arr = ", arr)
}
