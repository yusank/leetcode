/*
 * @lc app=leetcode.cn id=191 lang=golang
 *
 * [191] 位1的个数
 */

 package main

 import "log"

 func main() {
	log.Println(hammingWeight(8))
 }

// @lc code=start
func hammingWeight(num uint32) int {
    var (
		rst int
		pow = uint32(31)
	)

	for num != 0 {
		log.Println((num&1) << pow)
		if (num&1) << pow ==  1 {
			rst++
		}

		num = num >> 1
		pow--
	}

	return rst
}
// @lc code=end

