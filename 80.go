package main

import "fmt"

func main() {
	removeDuplicates80([]int{1, 1, 1, 2, 2, 3})
}

func removeDuplicates80(nums []int) int {
	n := len(nums)
	if n <= 2 {
		return n
	}
	fast, slow := 2, 2
	for fast < n {
		if nums[slow-2] != nums[fast] {
			// nums[0] nums[3]
			// 1 2
			// slow 的初始位置在 index 2
			nums[slow] = nums[fast]
			slow++
			fmt.Print(nums)
		}
		fast++
	}
	return slow
}
