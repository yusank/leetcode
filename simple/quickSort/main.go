package main

import "fmt"

// 快排

func main() {
	arr := []int{2, 3, 6, 1, 9, 8}
	quickSort(arr)
	fmt.Println(arr)
}

func quickSort(nums []int) {
	if len(nums) < 2 {
		return
	}

	var (
		tag        = nums[0]
		head, tail = 0, len(nums) - 1
	)

	for i := 1; i <= tail; {
		if nums[i] < tag {
			nums[i], nums[head] = nums[head], nums[i]
			i++
			head++
		} else {
			nums[i], nums[tail] = nums[tail], nums[i]
			tail--
		}
	}

	quickSort(nums[:head])
	quickSort(nums[head+1:])
}
