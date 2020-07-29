package leetcode
import(
    "testing"
)

func  TestTwoSum(t *testing.T){
  nums := []int{2,6,11,15}
  target := 9
  res := []int{0, 1}

  for i, v := range twoSum(nums, target) {
    if v != res[i] {
      t.Errorf("BinarySearch fail v is %d", v)
    }
  }
}

