package main

import(
    "fmt"
)

var(
    pre     []int
)

func UnionSearch(root int) int {
    son := root
    for root != pre[root] {
        root = pre[root]
    }
    for son != root {
        pre[son], son = root, pre[son]
    }
    return root
}

func Join(root1, root2 int) {
    x, y := UnionSearch(root1), UnionSearch(root2)
    if x != y {
        pre[x] = y
    }
}


func main() {
    edges := [][]int{[]int{1,2}, []int{2,3}, []int{3,4}, []int{4,1}, []int{1,5}}
    pre = make([]int, len(edges)+1)
    for i := range pre {
        pre[i] = i
    }
    for i := range edges {
        root1, root2 := UnionSearch(edges[i][0]), UnionSearch(edges[i][1])
        if root1 != root2 {
            // Join(root1, root2)
            pre[root1] = root2
        }else{
            fmt.Println("edges ", edges[i], "root1 ", root1, "root2 ", root2)
        }
        fmt.Println(pre)
    }
}
