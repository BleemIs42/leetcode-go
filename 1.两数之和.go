/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */
package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	ret := twoSum(nums, target)
	fmt.Println(ret)
}

// @lc code=start
func twoSum(nums []int, target int) []int {
	hashMap := make(map[int]int, len(nums))
	for k, v := range nums {
		rest := target - v
		idx, ok := hashMap[rest]
		if ok {
			return []int{idx, k}
		} else {
			hashMap[v] = k
		}
	}
	return []int{}
}

// @lc code=end
