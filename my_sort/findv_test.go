package mysort

import(
    "testing"
      "fmt"
    )

func TestFindV(t *testing.T){
  nums := []int{2,3,4,5,6,7,1,9,6,8}
  v := 6
  i := FindV(nums, v)
  if i != 4 {
    t.Errorf("6 in nums [2,3,4,5,6,7,1,9,6,8] is %d , is not 4!", i)
  }

}
func TestAdd(t *testing.T){
  nums1 := []int{1,0,1,1,1,1}
  nums2 := []int{1,0,1,1,1,1}
  numsc := Add(nums1, nums2)
           fmt.Println(numsc)
    if numsc[6] != 1{
      t.Errorf("The numsc[6] is not 1, is %d", numsc[6])
    }
}
