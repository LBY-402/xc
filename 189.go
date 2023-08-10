package main

func main() {
	rotate([]int{1, 2, 3, 4, 5, 6, 7}, 3)
}

// 注意 go 的Slice和 Map 是浅拷贝，其余（如 Array，Struct，Int, Bool, String）均为深拷贝
func rotate(nums []int, k int) {
	newNums := make([]int, len(nums))
	for i, v := range nums {
		newNums[(i+k)%len(nums)] = v
	}
	copy(nums, newNums)
}
