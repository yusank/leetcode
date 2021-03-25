/*
 * @lc app=leetcode.cn id=48 lang=golang
 *
 * [48] 旋转图像
 */
package main

func main() {

}

// @lc code=start
func rotate(matrix [][]int) {
	// 水平旋转
	n := len(matrix)
	for i := 0; i < n/2; i++ {
		matrix[i], matrix[n-1-i] = matrix[n-1-i], matrix[i]
	}

	// 对角线
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}

// @lc code=end
