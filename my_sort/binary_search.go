package mysort

import(

    )

func BinarySearch(nums []int, v int) int{
  start := 0
  end := len(nums)
  harf := end/2
  for start < harf && end > harf{
    if v == nums[harf]{
      return harf
    }
    if v < nums[harf]{
      end = harf 
      harf = harf/2
    }
    if v > nums[harf]{
      start = harf
      harf = (end - start)/2 +harf 
    }
  }
  if harf == 0{
    return -1
  }
  return harf
}
