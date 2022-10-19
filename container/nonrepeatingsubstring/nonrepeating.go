package main

import "fmt"

/*// for performance promotion, refactor map of variable lastOccurred to be as []int - slice, thus reduce elapsed time for assigning or accessing the map.
func lengthOfNonRepeatingSubStr(s string) int {
	// stores last occurred pos + 1, 0 means not seen
	//0x represents hexadecimal, 0xffff equals 65535 in decimal
	lastOccurred := make([]int, 0xffff) 	//via bench testing & cpu-profile analysis, this statement would lead to huge occupation and slow down the performance.
	start := 0
	maxlength := 0

	for i, r := range []rune(s) { //rune occupies 4 bytes, Chinese character occupies 2 bytes.
		if lastI := lastOccurred[r]; start <= lastI {
			start = lastI
		}

		if i-start+1 > maxlength {
			maxlength = i - start + 1
		}

		lastOccurred[r] = i + 1
	}

	return maxlength
}*/

// improve the performance after investigate causes of method above.
var lastOccurred = make([]int, 0xffff)

func lengthOfNonRepeatingSubStr(s string) int {
	for i := range lastOccurred {
		lastOccurred[i] = 0
	}

	start := 0
	maxlength := 0

	for i, r := range []rune(s) { //rune occupies 4 bytes, Chinese character occupies 2 bytes.
		if lastI := lastOccurred[r]; start <= lastI {
			start = lastI
		}

		if i-start+1 > maxlength {
			maxlength = i - start + 1
		}

		lastOccurred[r] = i + 1
	}

	return maxlength
}

//original method to count the max-length of non-repeating substring
func lengthOfNonrepeatingSubstrOriginal(s string) int {
	lastOccurred := make(map[rune]int)
	start := 0
	maxlength := 0

	for i, r := range []rune(s) {
		if lastI, ok := lastOccurred[r]; ok && start <= lastI {
			start = lastI + 1
		}

		if maxlength < i-start+1 {
			maxlength = i - start + 1
		}

		lastOccurred[r] = i
	}
	return maxlength
}

func main() {
	fmt.Println(lengthOfNonRepeatingSubStr("abcabcbb"))
	fmt.Println(lengthOfNonRepeatingSubStr("bbbbb"))
	fmt.Println(lengthOfNonRepeatingSubStr("pwwkew"))
	fmt.Println(lengthOfNonRepeatingSubStr(""))
	fmt.Println(lengthOfNonRepeatingSubStr("b"))
	fmt.Println(lengthOfNonRepeatingSubStr("abcdef"))
	fmt.Println(lengthOfNonRepeatingSubStr("一二三二一")) //需要用rune，才能得到正确的中文计算结果
	fmt.Println(lengthOfNonRepeatingSubStr("许洋洋"))
	fmt.Println(lengthOfNonRepeatingSubStr("黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花"))
}
