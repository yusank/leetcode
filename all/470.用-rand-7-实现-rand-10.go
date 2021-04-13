/*
 * @lc app=leetcode.cn id=470 lang=golang
 *
 * [470] 用 Rand7() 实现 Rand10()
 */

// @lc code=start
func rand10() int {
	// 1 2 3 4 5 6 7
	// 1 2 4 5 7 8 10
	// 1 2 3 4 5 6 7 8 9 10

	for {
		num := (rand7()-1)*7 + rand7()
		if num <= 40 {
			return 1 + num%10
		}

		num = (num-40-1)*7 + rand7()
		if num <= 60 {
			return 1+ num%10
		}

		num = (num-60-1)*7 + rand7()
		if num <= 20 {
			return 1+ num%10
		}
	}

	return -1
}
// @lc code=end

