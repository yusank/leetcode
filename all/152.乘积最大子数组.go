/*
 * @lc app=leetcode.cn id=152 lang=golang
 *
 * [152] 乘积最大子数组
 */

// @lc code=start
func maxProduct(nums []int) int {
	var (
		maxF, minF, ans = nums[0], nums[0], nums[0]
	)

	// 没太懂这个逻辑
	// 2 3 -2 4
	// maxF 维护一个从某个正数开始的正数结尾的最大乘积 越大越好
	// minF 维护一个从某个负数开始负数结尾的最小乘积 所以越小越好 最后负负得正
	// 遍历的时候 i 既是这个结尾 所以找 maxF 和 minF 中更大的一个
	// maxF 为例 对比三个值 num[i], 截止i-1的最大数maxF * nums[i] , 截止i-1的最小负数minF * nums[i]
	for i := 1; i < len(nums); i++ {
		mx, mm := maxF, minF
		maxF = max(mx*nums[i], max(nums[i], mm*nums[i]))
		minF = min(mm*nums[i], min(nums[i], mx*nums[i]))

		ans = max(ans, maxF)
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}

	return a
}

// @lc code=end

