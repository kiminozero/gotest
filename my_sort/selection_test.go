package mysort

import(
    "fmt"
    "testing"
    )

func TestSelectionSort(t *testing.T){
  nums := []int{2,3,5,6,1,0,9,8,7,4}
  SelectionSort(nums)
  fmt.Println("SelectionSort", nums)
}

