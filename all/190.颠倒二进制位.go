/*
 * @lc app=leetcode.cn id=190 lang=golang
 *
 * [190] 颠倒二进制位
 */

// @lc code=start
func reverseBits(num uint32) uint32 {
	var (
		res uint32
		pow = uint32(31)
	)

	for num != 0 {
		res += (num&1) << pow
		num = num >> 1
		pow--
	}

	return res
}
// @lc code=end

