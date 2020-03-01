package mysort

func FindV(nums []int, v int) int{
  for i, value := range(nums){
    if v == value{
      return i
    }
  }
  return -1
}

func Add(nums1, nums2 []int) []int{
    // two binary numbers array add
  numsc := make([]int, len(nums1)+1)
  temp := 0
  for i := 0; i < len(nums1); i++{
      numsc[i] = nums1[i]^nums2[i]^temp
      temp = nums1[i] + nums2[i] + temp
      temp = temp >> 1
  }
  numsc[len(nums1)]=temp
  return numsc
}
