package main

func main() {

}

// 贪心
func maxProfit122(prices []int) int {
	ans := 0
	for i := 1; i < len(prices); i++ {
		ans += max(0, prices[i]-prices[i-1])
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
