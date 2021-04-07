/*
 * @lc app=leetcode.cn id=167 lang=golang
 *
 * [167] 两数之和 II - 输入有序数组
 */

// @lc code=start
func twoSum(numbers []int, target int) []int {
	l,h := 0,len(numbers)-1
	for l<h {
		s := numbers[l] + numbers[h]
		if s == target {
			return []int{l+1,h+1}
		}

		if s < target {
			l++
		} else {
			h--
		}
	}

	return nil
}

// @lc code=end

