/*
 * @lc app=leetcode.cn id=104 lang=golang
 *
 * [104] 二叉树的最大深度
 *
 * https://leetcode-cn.com/problems/maximum-depth-of-binary-tree/description/
 *
 * algorithms
 * Easy (71.30%)
 * Likes:    565
 * Dislikes: 0
 * Total Accepted:    187.2K
 * Total Submissions: 255.1K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * 给定一个二叉树，找出其最大深度。
 *
 * 二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。
 *
 * 说明: 叶子节点是指没有子节点的节点。
 *
 * 示例：
 * 给定二叉树 [3,9,20,null,null,15,7]，
 *
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 *
 * 返回它的最大深度 3 。
 *
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	tree := TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:  2,
			Left: nil,
		},
		Right: &TreeNode{
			Val:  1,
			Left: nil,
		},
	}
	res := maxDepth(&tree)
	fmt.Println(res)
}

// @lc code=start
func maxDepth(root *TreeNode) int {
	return maxDepthQ(root)

	if root == nil {
		return 0
	}
	lh := maxDepth(root.Left)
	rh := maxDepth(root.Right)
	return maxInt(lh, rh) + 1
}
func maxInt(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func maxDepthQ(root *TreeNode) int {
	if root == nil {
		return 0
	}
	
	q := []*TreeNode{root}
	deep := 0
	for len(q) > 0 {
		deep++
		l := len(q)
		for i := 0; i < l; i++ {
			curr := q[i]
			if curr.Left != nil {
				q = append(q, curr.Left)
			}
			if curr.Right != nil {
				q = append(q, curr.Right)
			}
		}
		q = q[l:]
	}
	return deep
}

// @lc code=end
