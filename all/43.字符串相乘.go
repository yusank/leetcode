/*
 * @lc app=leetcode.cn id=43 lang=golang
 *
 * [43] 字符串相乘
 */

package main

import (
	"log"
	"strconv"
)

func main() {
	log.Println(multiply("123", "456") == "56088")
}

// @lc code=start
func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	m, n := len(num1), len(num2)
	slice := make([]int, m+n)

	for i := m - 1; i >= 0; i-- {
		x := int(num1[i] - '0')
		for j := n - 1; j >= 0; j-- {
			y := int(num2[j] - '0')
			slice[i+j+1] += x * y
		}
	}

	for i := n + m - 1; i > 0; i-- {
		slice[i-1] += slice[i] / 10
		slice[i] %= 10
	}

	var (
		str string
		idx int
	)

	if slice[0] == 0 {
		idx = 1
	}

	for ; idx < n+m; idx++ {
		str += strconv.Itoa(slice[idx])
	}

	return str
}

// @lc code=end
