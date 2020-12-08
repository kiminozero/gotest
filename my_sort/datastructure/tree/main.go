package main

import (
    "fmt"
)

func main() {
    obj := Constructor()
    //data := "1,2,#,#,3,#,#"
    data := "4,2,1,#,#,3,#,#,6,5,#,#,7,#,#"
    root := obj.deserialize(data)
    fmt.Println("root.Val = ", root.Val)
    fmt.Println("inorder morris")
    InorderMorris(root)
    
    fmt.Println("preorder morris")
    PreorderMorris(root)

    fmt.Println("inorder")
    Inorder(root)

    fmt.Println("serialize")
    fmt.Println(obj.serialize(root))
}



