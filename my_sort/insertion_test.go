package mysort

import(
    "fmt"
    "testing"
    )

func TestInsertionSort(t *testing.T){
  nums := []int{2,3,5,6,1,0,9,8,7,4}
  InsertionSort(nums)
  fmt.Println("InsertionSort", nums)
}

func TestReverseInsertionSort(t *testing.T){
  nums := []int{2,3,5,6,1,0,9,8,7,4}
  ReverseInsertionSort(nums)
  fmt.Println("ReverseInsertionSort", nums)
}

func TestReverseInsertionSortR(t *testing.T){
  nums := []int{3,2,5,6,1,0,9,8,7,4}
  InsertionSortR(nums, len(nums)-1)
  fmt.Println("InsertionSortR", nums)
}

func TestReverseInsertionSortB(t *testing.T){
  nums := []int{3,2,5,6,1,0,9,8,7,4}
  InsertionSortB(nums)
  fmt.Println("InsertionSortB", nums)
}
