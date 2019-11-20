package main

import (
	"fmt"
)

/*
给定一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？找出所有满足条件且不重复的三元组。

注意：答案中不可以包含重复的三元组。

例如, 给定数组 nums = [-1, 0, 1, 2, -1, -4]，

满足要求的三元组集合为：
[
  [-1, 0, 1],
  [-1, -1, 2]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/3sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	// nums := []int{0, 0, 0, 0}

	fmt.Println(threeSum(nums))
}

func threeSum(nums []int) [][]int {
	var result = make([][]int, 0)

	if len(nums) < 3 {
		return result
	}
	quickSort(nums)
	fmt.Println(nums)

	for i := 0; i < len(nums)-1; i++ {
		var (
			head, end = i + 1, len(nums) - 1
		)
		if nums[i] > 0 || nums[head]+nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

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
	ln := len(nums)
	if ln < 2 {
		return
	}

	var (
		cur       = nums[0]
		head, end = 0, ln - 1
	)

	for i := 1; i <= end; {
		if nums[i] < cur {
			// swap
			nums[head], nums[i] = nums[i], nums[head]
			head++
			i++
		} else {
			nums[end], nums[i] = nums[i], nums[end]
			end--
		}
	}

	quickSort(nums[:head])
	quickSort(nums[head+1:])
}
