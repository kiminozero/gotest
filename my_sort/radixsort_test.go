package mysort
import(
    "testing"
    "fmt"
    )

func TestLSDSort(t *testing.T){
  nums := []int{21,32,35,24,16,37,61,90,89,78}
  LSDSort(nums)
  fmt.Println("LSDSort ",nums)
}

