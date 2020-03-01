package mysort

import(
    "testing"
    "fmt"
)

func TestQuickSort(t *testing.T) {
    nums := []int{1,9,8,6,5,7,4,3,2,0}
    QuickSort(nums, 0, 9)
    for k, v := range nums {
        if nums[k] != v {
            t.Errorf("QuickSort is wrong")
        }
    }
    fmt.Println("QuickSort ", nums)
}

func TestRandomizedQuickSort(t *testing.T) {
    nums := []int{1,9,8,6,5,7,4,3,2,0}
    RandomizedQuickSort(nums, 0, 9)
    for k, v := range nums {
        if nums[k] != v {
            t.Errorf("RandomizedQuickSort is wrong")
        }
    }
    fmt.Println("RandomizedQuickSort ", nums)
}

func TestQuickSortS(t *testing.T) {
    nums := []int{1,9,8,6,5,7,4,3,2,0}
    QuickSortS(nums, 0, 9)
    for k, v := range nums {
        if nums[k] != v {
            t.Errorf("QuickSortS is wrong")
        }
    }
    fmt.Println("QuickSortS ", nums)
}

func TestRandomizedSelect(t *testing.T) {
    nums := []int{1,9,8,6,5,7,4,3,2,0}
    v := RandomizedSelect(nums, 0, 9, 3)
    if v != 2 {
        t.Errorf("RandomizedSelect is wrong")
    }
}

func TestRandomizedSelectI(t *testing.T) {
    nums := []int{1,9,8,6,5,7,4,3,2,0}
    v := RandomizedSelectI(nums, 0, 9, 3)
    if v != 2 {
        t.Errorf("RandomizedSelectI is wrong and it is %d", v)
    }
}
