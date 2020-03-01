package mysort
import(
    "fmt"
    )

func SelectionSort(nums []int){
  for i := 0; i<len(nums)-1; i++{
    key := i
    for j:=i+1; j<len(nums); j++{
      if nums[key] > nums[j]{
        key = j 
      }
    }
    nums[i],nums[key] = nums[key], nums[i]
  }
  fmt.Println(nums)
}

