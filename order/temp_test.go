package order

import (
	"fmt"
	"strconv"
	"testing"
)

func Test1(t *testing.T) {
	fmt.Println(myAtoi("    -42"))
}

func myAtoi(s string) int {
	// initial param
	intMin := -(1 << 31)
	intMax := 1<<31 - 1
	sum := 0
	flag := 1
	start := true

	// loop
	for _, bt := range s {
		if string(bt) == " " && start {
			continue
		}
		if (string(bt) == "-" || string(bt) == "+") && start {
			start = false
			if string(bt) == "-" {
				flag = -1
			}
			continue
		}

		if string(bt) <= "9" && string(bt) >= "0" {
			tail, _ := strconv.ParseInt(string(bt), 0, 64)
			if flag == 1 && (sum > intMax/10 || (sum == intMax/10 && tail > int64(intMax%10))) {
				return intMax
			}
			if flag == -1 && (-sum < intMin/10 || (-sum == intMin/10 && -tail < int64(intMin%10))) {
				return intMin
			}
			sum = sum*10 + int(tail)
		} else {
			return flag * sum
		}
	}

	return flag * sum
}
