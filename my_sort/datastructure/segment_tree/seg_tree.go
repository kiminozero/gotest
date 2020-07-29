package main

import (
    "fmt"
    )


func build_tree(arr []int, tree *[]int, node, start, end int) {
  if start==end {
    (*tree)[node] = arr[start]
    return
  }
  mid := (start+end)>>1
  left_node  := 2*node+1
  right_node := 2*node+2

  build_tree(arr, tree, left_node, start, mid)
  build_tree(arr, tree, right_node, mid+1, end)
  (*tree)[node] = (*tree)[left_node] + (*tree)[right_node]
}

func update_tree(arr []int, tree *[]int, node, start, end, idx, val int) {
  if start == end {
    arr[idx] = val
    (*tree)[node] = val
    return
  }

  mid := (start+end)>>1
  left_node  := 2*node+1
  right_node := 2*node+2
  
  if idx >= start && idx <= mid {
    update_tree(arr, tree, left_node, start, mid, idx, val)
  } else {
    update_tree(arr, tree, right_node, mid+1, end, idx, val)
  }
  (*tree)[node] = (*tree)[left_node] + (*tree)[right_node]
}

func query_tree(arr []int, tree *[]int, node, start, end, L, R int) int {
  if L > end || R < start {
    return 0
  }
  if L <= start && R >= end || start==end{
    return (*tree)[node]
  }


  mid := (start+end)>>1
  left_node  := 2*node+1
  right_node := 2*node+2
  sum_left  := query_tree(arr, tree, left_node, start, mid, L, R)
  sum_right := query_tree(arr, tree, right_node, mid+1, end, L, R)

  return sum_left+sum_right
}

func main() {

  arr := []int{1,3,5,7,9,11}
  size := 6
  tree := make([]int, 4*size)
  build_tree(arr, &tree, 0, 0, size-1)
  for i := range tree {
    fmt.Println(tree[i])
  }

  update_tree(arr, &tree, 0, 0, size-1, 4, 6)

  for i := range tree {
    fmt.Println(tree[i])
  }

  sum := query_tree(arr, &tree, 0, 0, size-1, 2, 5)
  fmt.Println("sum = ", sum)
}
