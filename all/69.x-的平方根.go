/*
 * @lc app=leetcode.cn id=69 lang=golang
 *
 * [69] x 的平方根
 */

// @lc code=start
func mySqrt(x int) int {
	var (
		ans = -1
		l,r = 0,x
	)

	for l <= r {
		mid := (l+r)/2
		if mid*mid <= x {
			ans=mid
			l=mid+1
		} else {
			r=mid-1
		}
	}

	return ans
}
// @lc code=end

