package main

import (
	"fmt"
)

func main() {
	fmt.Println(isMatch([]int64{1, 5, 2, 4, 3}))
}

func isMatch(x []int64) bool {
	temp := make([]int64, len(x))
	for index := range temp {
		temp[index] = 1
	}
	var res int64 = 0
	for i, num := range x {
		for j := 0; j <= i; j++ {
			if num > x[j] {
				temp[i] = max(temp[i], x[j]+1)
			}
		}
		res = max(res, temp[i])
	}

	fmt.Println(res)
	return false
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
