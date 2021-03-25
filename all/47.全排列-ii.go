/*
 * @lc app=leetcode.cn id=47 lang=golang
 *
 * [47] 全排列 II
 */

// @lc code=start
func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	var (
		dfs  func(int)
		n    = len(nums)
		vis  = make([]bool, len(nums))
		temp []int
		rs   [][]int
	)

	dfs = func(idx int) {
		if idx == n {
			rs = append(rs, append([]int{}, temp...))
			return
		}

		for i := range nums {
			if vis[i] || i > 0 && !vis[i-1] && nums[i] == nums[i-1] {
				continue
			}

			temp = append(temp, nums[i])
			vis[i] = true
			dfs(idx + 1)
			// 回溯
			vis[i] = false
			temp = temp[:len(temp)-1]
		}
	}

	dfs(0)
	return rs
}

// @lc code=end

