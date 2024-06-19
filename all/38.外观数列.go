/*
 * @lc app=leetcode.cn id=38 lang=golang
 *
 * [38] 外观数列
 */

package main

import "strconv"

// @lc code=start
func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}

	last := countAndSay(n - 1)
	cur := last[0]
	curCount := 1
	var result string
	for i := 1; i < len(last); i++ {
		if last[i] == cur {
			curCount++
			continue
		}
		result += strconv.Itoa(curCount) + string(cur)
		cur = last[i]
		curCount = 1
	}

	result += strconv.Itoa(curCount) + string(cur)

	return result
}

// @lc code=end
