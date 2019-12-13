package main

type TreeNode struct {
  Val int
  Left *TreeNode
  Right *TreeNode
}

func diameterOfBinaryH (root *TreeNode, currDiameter int, currrHighestDiameter *int) int {
  if root.Left == nil && root.Right == nil {
    if currDiameter > *currrHighestDiameter {
      *currrHighestDiameter = currDiameter
    }

    return 0
  }

  var leftDepth int
  var rightDepth int

  if root.Left != nil {
    leftDepth = diameterOfBinaryH(root.Left, currDiameter+1, currrHighestDiameter)
  }

  if root.Right != nil {
    rightDepth = diameterOfBinaryH(root.Right, leftDepth+1, currrHighestDiameter)
  }

  return int(math.Max(float64(leftDepth+1), float64(rightDepth+1)))

}

func diameterOfBinary(root *TreeNode) int {
  if root == nil {
    return 0
  }

  var result int

  diameterOfBinaryH(root, 0, &result)

  return result
}

func main() {

}