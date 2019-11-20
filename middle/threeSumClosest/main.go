package main

import (
	"fmt"
)

/*
给定一个包括 n 个整数的数组 nums 和 一个目标值 target。找出 nums 中的三个整数，使得它们的和与 target 最接近。返回这三个数的和。假定每组输入只存在唯一答案。

例如，给定数组 nums = [-1，2，1，-4], 和 target = 1.

与 target 最接近的三个数的和为 2. (-1 + 2 + 1 = 2).

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/3sum-closest
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {
	nums := []int{-1, 2, 1, -4}

	fmt.Println(threeSumClosest(nums, 1))
}
func threeSumClosest(nums []int, n int) int {
	// 1. 找到所有的组合
	// 2. 得到的结果和 n 做距离计算
	// 3. 计算组合或距离时，遇到等于 n 的结果直接返回

	quickSort(nums)
	var (
		minDis = -1
		result int
	)
	for i := 0; i < len(nums)-1; i++ {
		var head, end = i + 1, len(nums) - 1

		for head < end {
			sum := nums[i] + nums[head] + nums[end]
			if sum == n {
				return sum
			}

			if sum > n {
				end--
			}

			if sum < n {
				head++
			}

			d := twoNumDistance(sum, n)
			if minDis == -1 || d < minDis {
				minDis = d
				result = sum
			}
		}
	}

	return result
}

func twoNumDistance(a, b int) int {
	if a > b {
		return a - b
	}

	return b - a
}

func quickSort(nums []int) {
	ln := len(nums)
	if ln < 2 {
		return
	}

	var (
		tag       = nums[0]
		head, end = 0, ln - 1
	)

	for i := 1; i <= end; {
		if nums[i] < tag {
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
