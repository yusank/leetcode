/*
 * @lc app=leetcode.cn id=191 lang=golang
 *
 * [191] 位1的个数
 */
package main

import "log"

func main() {
	log.Println(hammingWeight(9))
}

// @lc code=start
func hammingWeight(num uint32) int {
    var (
		rs int
	)

	for num != 0 {
		rs++
		num = num & (num-1)
	}
	return rs
}
// @lc code=end

