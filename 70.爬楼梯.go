/*
 * @lc app=leetcode.cn id=70 lang=golang
 *
 * [70] 爬楼梯
 *
 * https://leetcode-cn.com/problems/climbing-stairs/description/
 *
 * algorithms
 * Easy (47.20%)
 * Likes:    766
 * Dislikes: 0
 * Total Accepted:    115.8K
 * Total Submissions: 244K
 * Testcase Example:  '2'
 *
 * 假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
 *
 * 每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
 *
 * 注意：给定 n 是一个正整数。
 *
 * 示例 1：
 *
 * 输入： 2
 * 输出： 2
 * 解释： 有两种方法可以爬到楼顶。
 * 1.  1 阶 + 1 阶
 * 2.  2 阶
 *
 * 示例 2：
 *
 * 输入： 3
 * 输出： 3
 * 解释： 有三种方法可以爬到楼顶。
 * 1.  1 阶 + 1 阶 + 1 阶
 * 2.  1 阶 + 2 阶
 * 3.  2 阶 + 1 阶
 *
 *
 */
package main

import "fmt"

func main() {

	fmt.Println(3, climbStairsDP(3))
}

// @lc code=start
func climbStairs(n int) int {
	cache := map[int]int{}
	return helper(n, cache)
	// return climbStairsDP(n)
}

func helper(n int, cache map[int]int) int {
	if n <= 2 {
		return n
	}
	v, ok := cache[n]
	if ok {
		return v
	}
	v = helper(n-1, cache) + helper(n-2, cache)
	cache[n] = v
	return v
}

func climbStairsDP(n int) int {
	if n <= 1 {
		return n
	}

	dpOne, dpTwo := 1, 2
	for i := 3; i <= n; i++ {
		dpOne, dpTwo = dpTwo, dpOne+dpTwo
	}
	return dpTwo
}

// @lc code=end
