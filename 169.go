package main

func main() {

}

// 摩尔投票法（Boyer–Moore majority vote algorithm），也被称作「多数投票法」，
// 算法解决的问题是：如何在任意多的候选人中（选票无序），选出获得票数最多的那个。
//
// 我们遍历投票数组，将当前票数最多的候选人与其获得的（抵消后）票数分别存储在 major 与 count 中。
// 当我们遍历下一个选票时，判断当前 count 是否为零：
//
//若 count == 0，代表当前 major 空缺，直接将当前候选人赋值给 major，并令 count++
//若 count != 0，代表当前 major 的票数未被完全抵消，因此令 count--，即使用当前候选人的票数抵消 major 的票数
//
func majorityElement(nums []int) int {
	major := 0
	count := 0

	for _, num := range nums {
		if count == 0 {
			major = num
		}
		if major == num {
			count++
		} else {
			count--
		}
	}
	return major
}
