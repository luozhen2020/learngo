package main

import (
	"fmt"
	"unicode/utf8"
)

func runeLengthOfNonRepeatingSubStr(s string) int {
	lastOccurred := make(map[rune]int)
	start := 0
	maxLength := 0

	for i, ch := range []rune(s) {
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}

	return maxLength
}

func main() {
	s := "Yes我爱许洋洋！" //UTF-8
	/*fmt.Println(len(s))*/
	/*fmt.Printf("%X\n", []byte(s))*/
	for _, b := range []byte(s) {
		fmt.Printf("%X ", b)
	}
	fmt.Println()

	for i, ch := range s { //ch is a rune
		fmt.Printf("(%d %X) ", i, ch)
	}

	fmt.Println("Rune count: ", utf8.RuneCountInString(s))

	bytes := []byte(s)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c ", ch)
	}
	fmt.Println()

	for i, ch := range []rune(s) {
		fmt.Printf("(%d %c) ", i, ch)
	}
	fmt.Println()

	fmt.Println(runeLengthOfNonRepeatingSubStr("abcabcbb"))
	fmt.Println(runeLengthOfNonRepeatingSubStr("bbbbb"))
	fmt.Println(runeLengthOfNonRepeatingSubStr("pwwkew"))
	fmt.Println(runeLengthOfNonRepeatingSubStr(""))
	fmt.Println(runeLengthOfNonRepeatingSubStr("b"))
	fmt.Println(runeLengthOfNonRepeatingSubStr("abcdef"))
	fmt.Println(runeLengthOfNonRepeatingSubStr("一二三二一")) //需要用rune，才能得到正确的中文计算结果
	fmt.Println(runeLengthOfNonRepeatingSubStr("许洋洋!"))
	fmt.Println(runeLengthOfNonRepeatingSubStr("黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花"))
}
