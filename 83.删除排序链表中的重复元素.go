/*
 * @lc app=leetcode.cn id=83 lang=golang
 *
 * [83] 删除排序链表中的重复元素
 *
 * https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list/description/
 *
 * algorithms
 * Easy (48.22%)
 * Likes:    319
 * Dislikes: 0
 * Total Accepted:    106.1K
 * Total Submissions: 210.3K
 * Testcase Example:  '[1,1,2]'
 *
 * 给定一个排序链表，删除所有重复的元素，使得每个元素只出现一次。
 *
 * 示例 1:
 *
 * 输入: 1->1->2
 * 输出: 1->2
 *
 *
 * 示例 2:
 *
 * 输入: 1->1->2->3->3
 * 输出: 1->2->3
 *
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	ll := ListNode{
		Val:  1,
		Next: nil,
	}
	ll.Next = &ListNode{
		Val:  1,
		Next: nil,
	}
	ll.Next.Next = &ListNode{
		Val:  2,
		Next: nil,
	}
	res := deleteDuplicates(&ll)
	fmt.Println(res)
}

// @lc code=start
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	currNode := head

	for currNode != nil && currNode.Next != nil {
		if currNode.Val == currNode.Next.Val {
			currNode.Next = currNode.Next.Next
		} else {
			currNode = currNode.Next
		}
	}
	return head
}

// @lc code=end
