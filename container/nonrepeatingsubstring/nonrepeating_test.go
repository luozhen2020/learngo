package main

import "testing"

// 测试
// 覆盖率测试 （控制台、终端）
func TestSubstr(t *testing.T) {
	tests := []struct {
		s   string
		ans int
	}{
		// normal cases
		{"abcabcbb", 3},
		{"pwwkew", 3},

		// edge cases
		{"", 0},
		{"b", 1},
		{"bbbbb", 1},
		{"abcabcabcd", 4},

		// Chinese support
		{"一二三二一", 3},
		{"我的女神是许洋洋", 7},
		{"黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花", 8},
	}

	for _, tt := range tests {
		actual := lengthOfNonRepeatingSubStr(tt.s)
		if actual != tt.ans {
			t.Errorf("Got %d for input: %s, expected %d", actual, tt.s, tt.ans)
		}
	}
}

// 性能测试 （控制台、终端）
func BenchmarkSubstr(b *testing.B) {
	s, ans := "黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花", 8

	for i := 0; i < 13; i++ {
		s = s + s
	}
	b.Logf("len(s) = %d", len(s))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		/*Conclusion: 使用map处理长度较短的字符串时，更加高效；使用slice在处理长度非常长的字符串时，更加高效*/
		actual := /*lengthOfNonrepeatingSubstrOriginal(s)*/ lengthOfNonRepeatingSubStr(s)
		if actual != ans {
			b.Errorf("Got %d for input: %s, expected %d", actual, s, ans)
		}
	}
}
