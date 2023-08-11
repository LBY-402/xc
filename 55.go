package main

func main() {

}

func canJump(nums []int) bool {
	// 初始化一个最大能跳跃的步
	k := 0
	for i := 0; i < len(nums); i++ {
		if i > k {
			return false
		}

		// i + nums[i]  表示可以跳的最大步
		k = max(k, i+nums[i])
	}
	return true
}
