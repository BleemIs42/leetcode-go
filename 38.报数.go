/*
 * @lc app=leetcode.cn id=38 lang=golang
 *
 * [38] 报数
 *
 * https://leetcode-cn.com/problems/count-and-say/description/
 *
 * algorithms
 * Easy (53.22%)
 * Likes:    350
 * Dislikes: 0
 * Total Accepted:    62.2K
 * Total Submissions: 116.1K
 * Testcase Example:  '1'
 *
 * 报数序列是一个整数序列，按照其中的整数的顺序进行报数，得到下一个数。其前五项如下：
 *
 * 1.     1
 * 2.     11
 * 3.     21
 * 4.     1211
 * 5.     111221
 *
 *
 * 1 被读作  "one 1"  ("一个一") , 即 11。
 * 11 被读作 "two 1s" ("两个一"）, 即 21。
 * 21 被读作 "one 2",  "one 1" （"一个二" ,  "一个一") , 即 1211。
 *
 * 给定一个正整数 n（1 ≤ n ≤ 30），输出报数序列的第 n 项。
 *
 * 注意：整数顺序将表示为一个字符串。
 *
 *
 *
 * 示例 1:
 *
 * 输入: 1
 * 输出: "1"
 *
 *
 * 示例 2:
 *
 * 输入: 4
 * 输出: "1211"
 *
 *
 */
package main

import "fmt"

import "strconv"

import "strings"

func main() {
	n := 5
	ret := countAndSay(n)
	fmt.Println(n, ret)
}

// @lc code=start
func countAndSay(n int) string {
	if n < 1 || n > 30 {
		fmt.Println("Please input 1 <= n <= 30")
		return "0"
	}

	if n == 1 {
		return "1"
	}

	arr := []int{1}
	for index := 2; index <= n; index++ {
		var tempArr []int
		count := 0
		temp := arr[0]
		for _, v := range arr {
			if temp == v {
				count++
			} else {
				tempArr = append(tempArr, count, temp)
				temp = v
				count = 1
			}
		}
		arr = append(tempArr, count, temp)
	}

	ret := []string{}
	for _, v := range arr {
		ret = append(ret, strconv.Itoa(v))
	}
	return strings.Join(ret, "")
}

// @lc code=end
