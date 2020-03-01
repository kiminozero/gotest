package mysort

import(
    "testing"
    "fmt"
    )

func TestHeapSort(t *testing.T){
  nums := []int{2,3,5,4,6,7,1,0,9,8}
  HeapSort(nums)
  fmt.Println("HeapSort ",nums)
}

func TestHeapSortReverse(t *testing.T){
  nums := []int{2,3,5,4,6,7,1,0,9,8}
  HeapSortReverse(nums)
  fmt.Println("HeapSortReverse ",nums)
}
