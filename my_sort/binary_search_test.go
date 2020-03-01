package mysort

import(
    "testing"
    )

func  TestBinarySearch(t *testing.T){
  nums := []int{0,1,2,3,4,5,6,7,8,9, 10}
  values := []int{1, 3, 4, 5, 7, 9, 10}
  for _, v := range(values){
    if v != BinarySearch(nums, v){
      t.Errorf("BinarySearch fail v is %d", v)
    }
  }

}
