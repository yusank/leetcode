/*
 * @lc app=leetcode.cn id=45 lang=golang
 *
 * [45] 跳跃游戏 II
 */

// [7,0,9,6,9,6,1,7,9,0,1,2,9,0,3]
// @lc code=start
func jump(nums []int) int {
	var (
		m   int
		cnt int
		end int
	)
	for i := 0; i < len(nums)-1; i++ {
		m = max(m, nums[i]+i)
		if i == end {
			end = m
			cnt++
		}
	}

	return cnt
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

// @lc code=end

