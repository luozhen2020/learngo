package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func getBalls() {
	// 定义红球数组
	var redBalls [6]int
	for i := 0; i < 6; i++ {
		for {
			// 取随机红球
			rand.Seed(time.Now().UnixNano())
			num := rand.Intn(33) + 1
			// 判断红球是否重复
			if ifInBalls(num, redBalls) {
				redBalls[i] = num
				break
			} else {
				continue
			}
		}
	}
	// 从小到大排序
	sort.Ints(redBalls[:])

	// 定义篮球数组
	var blueBall [1]int
	// 取随机篮球
	num := rand.Intn(16) + 1
	blueBall[0] = num

	// 输出结果
	fmt.Println(redBalls, blueBall)
}

// 验证红球是否重复func
func ifInBalls(param int, list [6]int) bool {
	for _, b := range list {
		if b == param {
			return false
		}
	}
	return true
}

func main() {
	getBalls()
}
