/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 */

// @lc code=start
func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}

	quickSort(nums)

	result := [][]int{}

	for i := 0; i < len(nums)-1; i++ {
		var (
			head, end = i + 1, len(nums) - 1
		)

		// 再往后三数之和大于0
		if nums[i] > 0 || nums[i]+nums[head] > 0 {
			break
		}

		// 重复
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		// head end 指针往后移动
		for head < end {
			if head > i+1 && nums[head] == nums[head-1] {
				head++
				continue
			}

			if end < len(nums)-2 && nums[end] == nums[end+1] {
				end--
				continue
			}

			sum := nums[i] + nums[head] + nums[end]

			if sum > 0 {
				end--
				continue
			}

			if sum < 0 {
				head++
				continue
			}

			result = append(result, []int{nums[i], nums[head], nums[end]})
			head++
			end--
		}
	}

	return result
}

func quickSort(nums []int) {
	n := len(nums)
	if n < 2 {
		return
	}

	var (
		cur       = nums[0]
		head, end = 0, n - 1
	)

	for i := 1; i <= end; {
		if nums[i] < cur {
			nums[i], nums[head] = nums[head], nums[i]
			head++
			i++
		} else {
			nums[i], nums[end] = nums[end], nums[i]
			end--
		}
	}

	quickSort(nums[:head])
	quickSort(nums[head+1:])
}

// @lc code=end

