/*
 * @lc app=leetcode.cn id=118 lang=golang
 *
 * [118] 杨辉三角
 *
 * https://leetcode-cn.com/problems/pascals-triangle/description/
 *
 * algorithms
 * Easy (64.72%)
 * Likes:    313
 * Dislikes: 0
 * Total Accepted:    85K
 * Total Submissions: 127.5K
 * Testcase Example:  '5'
 *
 * 给定一个非负整数 numRows，生成杨辉三角的前 numRows 行。
 *
 *
 *
 * 在杨辉三角中，每个数是它左上方和右上方的数的和。
 *
 * 示例:
 *
 * 输入: 5
 * 输出:
 * [
 * ⁠    [1],
 * ⁠   [1,1],
 * ⁠  [1,2,1],
 * ⁠ [1,3,3,1],
 * ⁠[1,4,6,4,1]
 * ]
 *
 */
package main

import "fmt"

func main() {
	res := generate(5)
	fmt.Println(res)
}

// @lc code=start
func generate(numRows int) [][]int {
	var res [][]int
	if numRows == 0 {
		return res
	}
	res = append(res, []int{1})
	for i := 1; i < numRows; i++ {
		row := []int{0}
		row = append(row, res[i-1]...)
		for j := 0; j < len(row)-1; j++ {
			row[j] = row[j] + row[j+1]
		}
		res = append(res, row)
	}
	return res
}

// @lc code=end
