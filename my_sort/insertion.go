package mysort
import(
    "fmt"
    )

func InsertionSort(nums []int){
  for j := 1; j<len(nums); j++{
    key := nums[j]
    // Insert nums[j] into the sorted sequence nums[0..j-1]
    i := j-1
    for i>-1 && nums[i] > key{
      nums[i + 1] = nums[i]
      i--
    }
    nums[i+1] = key
  }
  fmt.Println(nums)
}

func ReverseInsertionSort(nums []int){
  for j := 1; j<len(nums); j++{
    key := nums[j]
    // Insert nums[j] into the sorted sequence nums[0..j-1]
    i := j-1
    for i>-1 && nums[i] < key{
      nums[i + 1] = nums[i]
      i--
    }
    nums[i+1] = key
  }
  fmt.Println(nums)
}


func InsertionSortR(nums []int, end int){
  if end > 1{
    InsertionSortR(nums, end-1)
    SortR(nums, end)
  } 
  SortR(nums, end)
}

func SortR(nums []int, end int){
  key := nums[end]
  i := end-1
  for i > -1 &&nums[i] > key{
    nums[i+1] =nums[i]
    i--
  }
  nums[i+1] = key
}

func InsertionSortB(nums []int){
  for j := 1; j<len(nums); j++{
    key := nums[j]
    // Insert nums[j] into the sorted sequence nums[0..j-1]
    harf := BinarySearch(nums[:j], key)
    i := j-1
    // fmt.Println("harf", harf, "key",key, "nums", nums, nums[:j])
    for i > harf{
      nums[i + 1] = nums[i]
      i--
    }     
    nums[i+1] = key
  }
  fmt.Println(nums)
}

