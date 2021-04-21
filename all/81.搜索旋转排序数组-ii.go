/*
 * @lc app=leetcode.cn id=81 lang=golang
 *
 * [81] 搜索旋转排序数组 II
 */

// @lc code=start
func search(nums []int, target int) bool {
	var (
		l, r = 0, len(nums) - 1
		n    = len(nums)
	)

	if n == 0 {
		return false
	}

	if n == 1 {
		return nums[0] == target
	}

	for l < r {
		if nums[l] == target || nums[r] == target {
			return true
		}
		mid := (l + r) / 2
		if nums[mid] == target {
			return true
		}

		if nums[l] == nums[mid] && nums[mid] == nums[r] {
			l++
			r--
			continue
		}

		// 左边有序
		fmt.Println(l, r, mid)
		if nums[l] <= nums[mid] {
			if nums[mid] > target && nums[l] < target {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			if nums[mid] < target && nums[r] > target {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}

	return false
}

// @lc code=end

