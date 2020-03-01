package mysort

import(
    "testing"
    "fmt"
)

func TestCountingSort(t *testing.T) {
    nums := []int{1,0,2,4,5,7,8,9,0,3}
    want := []int{0,0,1,2,3,4,5,7,8,9}
    k := 10
    B := make([]int, len(nums))
    CountingSort(nums, B, k)

    for i := 0; i<len(nums); i++ {
        if B[i] != want[i] {
            t.Errorf("B[%d], %d != A[%d], %d ", i, B[i], i, want[i])
        }
    }
    fmt.Println(B)
}
