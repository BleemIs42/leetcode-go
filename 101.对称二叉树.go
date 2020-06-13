/*
 * @lc app=leetcode.cn id=101 lang=golang
 *
 * [101] 对称二叉树
 *
 * https://leetcode-cn.com/problems/symmetric-tree/description/
 *
 * algorithms
 * Easy (49.06%)
 * Likes:    846
 * Dislikes: 0
 * Total Accepted:    160.4K
 * Total Submissions: 307.4K
 * Testcase Example:  '[1,2,2,3,4,4,3]'
 *
 * 给定一个二叉树，检查它是否是镜像对称的。
 *
 * 例如，二叉树 [1,2,2,3,4,4,3] 是对称的。
 *
 * ⁠   1
 * ⁠  / \
 * ⁠ 2   2
 * ⁠/ \ / \
 * 3  4 4  3
 *
 * 但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的:
 *
 * ⁠   1
 * ⁠  / \
 * ⁠ 2   2
 * ⁠  \   \
 * ⁠  3    3
 *
 * 进阶：
 *
 * 你可以运用递归和迭代两种方法解决这个问题吗？
 *
 * Definition for a binary tree node.
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
	res := isSymmetric(&tree)
	fmt.Println(res)
}

// @lc code=start
func isSymmetric(root *TreeNode) bool {
	return isSymmetricQ(root)
	if root == nil {
		return true
	}
	return check(root.Left, root.Right)
}

func check(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	return p.Val == q.Val &&
		check(p.Left, q.Right) &&
		check(p.Right, q.Left)
}

func isSymmetricQ(root *TreeNode) bool {
	if root == nil {
		return true
	}
	q := []*TreeNode{root.Left, root.Right}
	for len(q) > 0 {
		left, right := q[0], q[1]
		q = q[2:]
		if left == nil && right == nil {
			continue
		}
		if left == nil || right == nil {
			return false
		}
		if left.Val != right.Val {
			return false
		}
		q = append(q, left.Left, right.Right)
		q = append(q, left.Right, right.Left)
	}
	return true
}

// @lc code=end
