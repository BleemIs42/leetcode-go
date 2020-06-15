/*
 * @lc app=leetcode.cn id=108 lang=golang
 *
 * [108] 将有序数组转换为二叉搜索树
 *
 * https://leetcode-cn.com/problems/convert-sorted-array-to-binary-search-tree/description/
 *
 * algorithms
 * Easy (67.60%)
 * Likes:    434
 * Dislikes: 0
 * Total Accepted:    72.6K
 * Total Submissions: 101.8K
 * Testcase Example:  '[-10,-3,0,5,9]'
 *
 * 将一个按照升序排列的有序数组，转换为一棵高度平衡二叉搜索树。
 *
 * 本题中，一个高度平衡二叉树是指一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1。
 *
 * 示例:
 *
 * 给定有序数组: [-10,-3,0,5,9],
 *
 * 一个可能的答案是：[0,-3,9,-10,null,5]，它可以表示下面这个高度平衡二叉搜索树：
 *
 * ⁠     0
 * ⁠    / \
 * ⁠  -3   9
 * ⁠  /   /
 * ⁠-10  5
 *
 *
 */

package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	list := []int{-10, -3, 0, 5, 9}
	res := sortedArrayToBST(list)
	fmt.Println(res)
}

// @lc code=start
func sortedArrayToBST(nums []int) *TreeNode {
	return sortedArrayToBSTQ(nums)

	if len(nums) == 0 {
		return nil
	}
	mid := len(nums) / 2

	root := &TreeNode{Val: nums[mid], Left: nil, Right: nil}
	root.Left = sortedArrayToBST(nums[:mid])
	root.Right = sortedArrayToBST(nums[mid+1:])

	return root
}

type TreeNodeQ struct {
	Node  *TreeNode
	Start int
	End   int
}

func sortedArrayToBSTQ(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	l := len(nums)
	root := &TreeNode{}
	queue := []TreeNodeQ{TreeNodeQ{Node: root, Start: 0, End: l}}
	for len(queue) > 0 {
		item := queue[0]
		start := item.Start
		end := item.End
		mid := (start + end) / 2
		item.Node.Val = nums[mid]
		if start < mid {
			item.Node.Left = &TreeNode{}
			queue = append(queue, TreeNodeQ{Node: item.Node.Left, Start: start, End: mid})
		}
		if mid+1 < end {
			item.Node.Right = &TreeNode{}
			queue = append(queue, TreeNodeQ{Node: item.Node.Right, Start: mid + 1, End: end})
		}
		queue = queue[1:]
	}
	return root
}

// @lc code=end
