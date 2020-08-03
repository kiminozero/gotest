package main

import(
    "strconv"
    "strings"
    "fmt"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
    Val     int
    Left    *TreeNode
    Right   *TreeNode
}

type Codec struct {

}

func Constructor() Codec {
    return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
    if root == nil {
        return "#"
    }
    l := this.serialize(root.Left)
    r := this.serialize(root.Right)
    return strconv.Itoa(root.Val) + "," + l  + "," + r
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
    if len(data) == 0 {
        return nil
    }
    nums := strings.Split(data, ",")

    var helper func() *TreeNode
    helper = func() *TreeNode {
        node := nums[0]
        nums = nums[1:]
        if len(nums) == 0 || node== "#" {
            return nil
        }
        v, _ := strconv.Atoi(node)
        root := &TreeNode{v, nil, nil}
        root.Left = helper()
        root.Right = helper()
        return root
    }
    // fmt.Println(data)
    return helper()
}


/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * data := obj.serialize(root);
 * ans := obj.deserialize(data);
 */

func main() {
    obj := Constructor()
    data := "1,2,#,#,3,#,#"
    root := obj.deserialize(data)
    fmt.Println(root.Val)
    fmt.Println(obj.serialize(root))
}
