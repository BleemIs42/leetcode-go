/*
 * @lc app=leetcode.cn id=67 lang=golang
 *
 * [67] 二进制求和
 *
 * https://leetcode-cn.com/problems/add-binary/description/
 *
 * algorithms
 * Easy (51.28%)
 * Likes:    277
 * Dislikes: 0
 * Total Accepted:    54.2K
 * Total Submissions: 105K
 * Testcase Example:  '"11"\n"1"'
 *
 * 给定两个二进制字符串，返回他们的和（用二进制表示）。
 *
 * 输入为非空字符串且只包含数字 1 和 0。
 *
 * 示例 1:
 *
 * 输入: a = "11", b = "1"
 * 输出: "100"
 *
 * 示例 2:
 *
 * 输入: a = "1010", b = "1011"
 * 输出: "10101"
 *
 */
package main

import "fmt"
import "strconv"

func main() {
	a := "1010"
	b := "1011"
	ret := addBinary(a, b)
	fmt.Println(a, b, ret)
}

// @lc code=start
func addBinary(a string, b string) string {
	i := len(a) - 1
	j := len(b) - 1

	ret := ""
	carry := 0

	for i >= 0 || j >= 0 {
		aVal, bVal := 0, 0
		if i >= 0 {
			aVal = int(a[i] - '0')
		}
		if j >= 0 {
			bVal = int(b[j] - '0')
		}
		sum := aVal + bVal + carry
		carry = 0
		if sum > 1 {
			carry = 1
		}
		ret = strconv.Itoa(sum%2) + ret

		i--
		j--
	}
	if carry > 0 {
		ret = "1" + ret
	}

	return ret
}

// @lc code=end
