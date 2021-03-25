/*
 * @lc app=leetcode.cn id=415 lang=golang
 *
 * [415] 字符串相加
 */

package main

import (
	"log"
	"strconv"
)

func main() {
	a := "998"
	b := "3"
	c := "1001"

	log.Println(addStrings(a, b), c)
}

// @lc code=start
func addStrings(num1 string, num2 string) string {
	var (
		maxStr, minStr = num1, num2
		gap            = len(maxStr) - len(minStr)
		rs             string
	)

	if len(maxStr) < len(minStr) {
		maxStr, minStr = minStr, maxStr
		gap = -gap
	}

	d := 0

	for i := len(maxStr) - 1; i >= 0; i-- {
		b1 := maxStr[i]
		b2 := byte('0')
		if i-gap >= 0 {
			b2 = minStr[i-gap]
		}

		sum := int(b1-'0') + int(b2-'0') + d
		d = sum / 10
		sum %= 10

		rs = strconv.Itoa(sum) + rs
	}

	if d > 0 {
		rs = strconv.Itoa(d) + rs
	}

	return rs

}

// @lc code=end
