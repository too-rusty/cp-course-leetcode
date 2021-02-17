/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {
    
    if root == nil { return nil }
    
    root.Left = invertTree(root.Left)
    root.Right = invertTree(root.Right)
    
    var tmp *TreeNode = root.Left
    root.Left = root.Right
    root.Right = tmp
    return root
    
}


/*


recursively invert the left subtree and the right subtree

and then swap them , this is the ease of recursion and its power

solve the sub problems and then merge them together to solve the bigger problem


*/
