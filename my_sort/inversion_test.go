package mysort
import(
    "testing"
    "fmt"
    )

func TestInversion(t *testing.T){
  nums := []int{2,3,8,6,1}
  count := 0
  Inversion(nums, 0, len(nums)-1, &count)
  fmt.Println("inversion", count)
}
