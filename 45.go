package main

func main() {

}

func jump(nums []int) int {
	end := 0
	l := len(nums)
	maxP := 0
	step := 0

	for i := 0; i < l-1; i++ {
		// 取最大的步数
		maxP = max(maxP, i+nums[i])
		// 更新边界
		if i == end {
			end = maxP
			step++
		}
	}
	return step
}
