package main

import "fmt"

func main() {
	fmt.Println(letterCombinations("223"))
}

func letterCombinations(digits string) []string {
	if digits == "" {
		return nil
	}

	var result []string
	for _, r := range digits {
		sli := numMap[string(r)]
		if len(sli) == 0 {
			continue
		}

		if len(result) == 0 {
			result = append(result, sli...)
			continue
		}

		var temp []string
		for _, item := range result {
			for _, s := range sli {
				temp = append(temp, item+s)
			}
		}
		result = temp
	}

	return result
}

var numMap = map[string][]string{
	"2": {"a", "b", "c"},
	"3": {"e", "d", "f"},
	"4": {"g", "h", "i"},
	"5": {"j", "k", "l"},
	"6": {"m", "n", "o"},
	"7": {"p", "q", "r", "s"},
	"8": {"t", "u", "v"},
	"9": {"w", "x", "y", "z"},
}
