/*
 * @lc app=leetcode.cn id=718 lang=golang
 *
 * [718] 最长重复子数组
 */

// @lc code=start
func findLength(nums1 []int, nums2 []int) (ans int) {
	return findMax(nums1, nums2)
	var (
		n  = len(nums1)
		m  = len(nums2)
		dp = make([][]int, n+1)
	)

	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, m+1)
	}

	for i := n - 1; i >= 0; i-- {
		for j := m - 1; j >= 0; j-- {
			if nums1[i] == nums2[j] {
				dp[i][j] = dp[i+1][j+1] + 1
			}

			if ans < dp[i][j] {
				ans = dp[i][j]
			}
		}
	}

	return
}

// 滑动窗口

func findMax(a, b []int) (ans int) {
	n, m := len(a), len(b)

	for i := 0; i < n; i++ {
		ln := min(m, n-i)
		ml := maxLen(a, b, i, 0, ln)
		ans = max(ans, ml)
	}

	for i := 0; i < m; i++ {
		ln := min(n, m-i)
		ml := maxLen(a, b, 0, i, ln)
		ans = max(ans, ml)
	}

	return
}

func min(a, b int) int {
	if a > b {
		return b
	}

	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func maxLen(a, b []int, i, j, ln int) int {
	var (
		cnt, m int
	)

	for k := 0; k < ln; k++ {
		if a[k+i] == b[j+k] {
			cnt++
			if cnt > m {
				m = cnt
			}
		} else {
			cnt = 0
		}
	}

	return m
}

// @lc code=end

