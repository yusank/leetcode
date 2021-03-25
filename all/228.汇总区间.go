/*
 * @lc app=leetcode.cn id=228 lang=golang
 *
 * [228] 汇总区间
 */

// @lc code=start
func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return nil
	}

	var (
		s, t   = 0, 0
		result []string
	)

	for i := 1; i < len(nums); i++ {
		if nums[i]-1 == nums[i-1] {
			t++
			continue
		}

		str := fmt.Sprintf("%d->%d", nums[s], nums[i-1])
		if s == t {
			str = fmt.Sprintf("%d", nums[s])
		}

		result = append(result, str)
		s = i
		t++
	}

	// todo 修改
	str := fmt.Sprintf("%d->%d", nums[s], nums[t])
	if s == t {
		str = fmt.Sprintf("%d", nums[s])
	}
	result = append(result, str)
	return result
}

// @lc code=end

