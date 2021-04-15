/*
 * @lc app=leetcode.cn id=74 lang=golang
 *
 * [74] 搜索二维矩阵
 */
package main

import "fmt"

func main() {
	matrix := [][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 60},
		// {1},
	}

	fmt.Println(searchMatrix(matrix, 7))
}

// @lc code=start
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}

	var (
		l int
		r = len(matrix) - 1
	)

	for l <= r {
		mid := l + (r-l)/2
		// fmt.Println(mid)

		if matrix[mid][0] <= target && matrix[mid][len(matrix[mid])-1] >= target {
			return findFromArray(matrix[mid], target)
		}

		if matrix[mid][0] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return false
}

func findFromArray(arr []int, target int) bool {
	if len(arr) == 0 {
		return false
	}

	// fmt.Println(arr, target)
	var (
		l int
		r = len(arr) - 1
	)

	for l <= r {
		mid := l + (r-l)/2
		// fmt.Println("arr:", l, r, mid)
		sub := arr[mid] - target
		if sub == 0 {
			return true
		}

		if sub > 0 {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return false
}

// @lc code=end
