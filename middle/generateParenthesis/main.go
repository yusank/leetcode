package main

import "fmt"
/*
给出 n 代表生成括号的对数，请你写出一个函数，使其能够生成所有可能的并且有效的括号组合。

例如，给出 n = 3，生成结果为：

[
  "((()))",
  "(()())",
  "(())()",
  "()(())",
  "()()()"
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/generate-parentheses
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {
	sli:= generateParenthesis(3)
	for i,v := range sli {
		fmt.Println(i,v)
	}
}

func generateParenthesis(n int) []string {
	if n ==0 {
		return nil
	}

	// 一组括号
	var result = make([][]string,n+1)
	result[0] = []string{""}
	result[1] = []string{"()"}

	// 从2 开始
	for i := 2; i < n+1; i++ {
		for j := 0; j < i; j++ {
			p := result[j]
			
			q := result[i-1-j]
			fmt.Println("---",i, j,p,q)
			for _, str1 := range result[j] {
				for _, str2 := range result[i-1-j] {

					el := "("+str1+")"+str2
					
					result[i] = append(result[i],el)
				}
			}
		}
	}

	return result[n]
} 