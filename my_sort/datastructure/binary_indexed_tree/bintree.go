package main

import(
    "fmt"
)

var (
    nums, tree      []int
)

func lowbit(x int) int {
    return x&(-x)
}

func add(i, v int) {
    for i <= len(nums) {
        tree[i] += v
        i += lowbit(i)
    }
}
func sum(i int) int {
    res := 0
    for i > 0 {
        res += tree[i]
        i -= lowbit(i)
    }
    return res
}

func sumFromTo(i, j int) int {
    return sum(j)-sum(i-1)
}

func main() {
    nums = []int{0,1,3,5,7,9, 11}
    tree = make([]int, len(nums)*2)
    for i := 1; i < len(nums); i++ {
        add(i, nums[i])
    }
    fmt.Println(nums, tree)
    fmt.Println("sum 2-5", sumFromTo(2, 5))
    fmt.Println("update 3 with 10")
    add(3, 10)
    fmt.Println(nums, tree)
    fmt.Println("sum 2-5", sumFromTo(2, 5))
}
