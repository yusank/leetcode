/*
 * @lc app=leetcode.cn id=67 lang=golang
 *
 * [67] 二进制求和
 */
 package main

 import "log"

 func main() {
	 a := "1001"
	 b := "11"

	 log.Println(addBinary(a,b))
 }

// @lc code=start
func addBinary(a string, b string) string {
    ans := ""
    carry := 0
    lenA, lenB := len(a), len(b)
    n := max(lenA, lenB)

    for i := 0; i < n; i++ {
        if i < lenA {
            carry += int(a[lenA-i-1] - '0')
        }
        if i < lenB {
            carry += int(b[lenB-i-1] - '0')
        }
        ans = bin2Str(carry) + ans
        carry /= 2
    }
    if carry > 0 {
        ans = "1" + ans
    }
    return ans
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}

func bin2Str(i int) string {
	if i % 2 == 0 {
		return "0"
	}

	return "1"
}

// @lc code=end

