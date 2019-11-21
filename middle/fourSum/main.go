package main

import (
	"fmt"
)

/*
给定一个包含 n 个整数的数组 nums 和一个目标值 target，判断 nums 中是否存在四个元素 a，b，c 和 d ，使得 a + b + c + d 的值与 target 相等？找出所有满足条件且不重复的四元组。

注意：

答案中不可以包含重复的四元组。

示例：

给定数组 nums = [1, 0, -1, 0, -2, 2]，和 target = 0。

满足要求的四元组集合为：
[
  [-1,  0, 0, 1],
  [-2, -1, 1, 2],
  [-2,  0, 0, 2]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/4sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {
	nums := []int{1, -2, -5, -4, -3, 3, 3, 5}
	target := -11

	// nums := []int{-1, -5, -5, -3, 2, 5, 0, 4}
	// target := -7

	// nums := []int{-3, -1, 0, 2, 4, 5}
	// target := 0

	// nums := []int{-1, 0, 1, 2, -1, -4}
	// target := -1

	fmt.Println(fourSum(nums, target))

}

func fourSum(nums []int, n int) [][]int {
	ln := len(nums)
	if ln < 4 {
		return nil
	}

	quickSort(nums)
	fmt.Println(nums)

	var result = make([][]int, 0)

	for i := 0; i < len(nums)-2; i++ {
		if n > 0 && nums[i] > n {
			break
		}

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for j := i + 1; j < len(nums)-1; j++ {
			var (
				head, end = j + 1, len(nums) - 1
			)
			fmt.Println(i, j, head, end)
			// if nums[j] > n || nums[head]+nums[j] > n {
			// 	break
			// }
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			for head < end {
				if head > j+1 && nums[head] == nums[head-1] {
					head++
					continue
				}

				if end < len(nums)-2 && nums[end] == nums[end+1] {
					end--
					continue
				}

				sum := nums[i] + nums[j] + nums[head] + nums[end]
				if sum > n {
					end--
					continue
				}

				if sum < n {
					head++
					continue
				}

				result = append(result, []int{nums[i], nums[j], nums[head], nums[end]})
				head++
				end--
			}
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
		tag        = nums[0]
		head, tail = 0, ln - 1
	)

	for i := 1; i <= tail; {
		if nums[i] < tag {
			nums[head], nums[i] = nums[i], nums[head]
			head++
			i++

			continue
		}

		nums[tail], nums[i] = nums[i], nums[tail]
		tail--
	}

	quickSort(nums[:head])
	quickSort(nums[head+1:])
}
