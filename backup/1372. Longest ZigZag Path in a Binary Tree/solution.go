/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func descendZigZag(node *TreeNode, goLeft bool, length int) int {
    if node == nil {
        return length-1
    }

    if node.Left == nil && node.Right == nil {
        return length
    }

    var l, r int
    if goLeft {
        l = descendZigZag(node.Left, false, length+1)
        r = descendZigZag(node.Right, true, 1)
    } else {
        l = descendZigZag(node.Left, false, 1)
        r = descendZigZag(node.Right, true, length+1)
    }

    if l > r {
        return l
    }

    return r
}

func longestZigZag(root *TreeNode) int {
    l := descendZigZag(root, false, 0)
    r := descendZigZag(root, true, 0)

    if l > r {
        return l
    }

    return r
}
