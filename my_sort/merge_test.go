package  mysort 

import(
    "fmt"
    "testing"
)

func TestMergeSort(t *testing.T){
    //nums := []int{3,4,2,1,5,6,7,8}
    nums := []int{9,3,4,2,1,5,6,7,8}
    MergeSort(nums, 0 ,len(nums)-1)
    fmt.Println(nums)

}

