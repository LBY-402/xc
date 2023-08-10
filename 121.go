package main

import "fmt"

/*



 */

func main() {
	fmt.Print(maxProfit([]int{2, 1, 2, 1, 0, 1, 2}))
}

func maxProfit(prices []int) int {
	fast, slow := 1, 0
	n := len(prices)
	count := 0
	for fast < n {
		if prices[fast] < prices[slow] {
			slow = fast
		} else {
			i := prices[fast] - prices[slow]
			if i > count {
				count = i
			}
		}

		fast++
	}
	return count
}
