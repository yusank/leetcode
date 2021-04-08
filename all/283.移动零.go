/*
 * @lc app=leetcode.cn id=283 lang=golang
 *
 * [283] 移动零
 */

// @lc code=start
func moveZeroes(nums []int) {
	// [0,1,2,0,0,3] 0 1
	// [1,0,2,0,0,3] 1 2
	// [1,2,0,0,0,3] 2 3
	// [1,2,3,0,0,0] 3 4

	var (
		start, end = -1, -1
	)

	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			if start < 0 {
				start = i
			}

			if end < 0 {
				end = i + 1
			} else {
				end++
			}

			continue
		}

		// 没有0
		if start == end {
			continue
		}

		nums[i], nums[start] = nums[start], nums[i]
		start++
		end++
	}

	return
}

// @lc code=end

