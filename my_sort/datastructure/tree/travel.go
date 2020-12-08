package main

import(
    "fmt"
)

func visit(root *TreeNode) {
    fmt.Println(root.Val)
}

func InorderMorris(root *TreeNode) {
    var cur, predecessor *TreeNode
    helper := func() {
        fmt.Println(cur.Val)
    }
    
    cur = root

    for cur != nil {

        if cur.Left != nil {
            predecessor = cur.Left 
            for predecessor.Right != nil && predecessor.Right != cur {
                predecessor = predecessor.Right
            }

            if predecessor.Right == nil {
                predecessor.Right = cur
                cur = cur.Left
            } else {
                helper()
                cur = cur.Right
                predecessor.Right = nil
            }
        } else {
            helper()
            cur = cur.Right
        }
    }
}


func PreorderMorris(root *TreeNode) {

    var cur, predecessor *TreeNode

    cur = root

    for cur != nil {
        if cur.Left != nil {
            predecessor = cur.Left
            for predecessor.Right != nil && predecessor.Right != cur {
                predecessor = predecessor.Right
            }

            if predecessor.Right == nil {
                visit(cur)
                predecessor.Right = cur
                cur = cur.Left
            } else {
                cur = cur.Right
                predecessor.Right = nil
            }
        } else {
            visit(cur)
            cur = cur.Right
        }
    }
}

func Inorder(root *TreeNode) {
    if root == nil {
        return
    }
    Inorder(root.Left)
    visit(root)
    Inorder(root.Right)
}
