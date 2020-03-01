package mysort

import(
    "fmt"
    )

func HeapMaximum(A []int) int{
  return A[0]
}

func HeapExtractMax(A []int) (int, error) {
  if len(A)< 1{
    return 0, fmt.Errorf("heap underflow")
  }
  max := A[0]
  A[0] = A[len(A)-1]
  maxHeapify(A, len(A) - 1, 0)
  
  return max, nil
}

func HeapIncreaseKey(A []int, i, key int) error{
  if key < A[i]{
    return fmt.Errorf("new key is smaller than current key")
  }
  A[i] = key
  for i > 0 && A[(i-1)/2] < A[i] {
    A[i], A[(i-1)/2] = A[(i-1)/2], A[i]
    i = (i-1)/2
  }
  return nil 
}

func MaxHeapInsert(A *[]int, key int){
  heapsize := len(*A) + 1
  *A = append(*A,INT_MIN)
  HeapIncreaseKey(*A, heapsize - 1, key)
}

func MaxHeapDelete(A *[]int, i int) error{
  if len(*A) < i + 1 {
    return fmt.Errorf("i not in heap")
  }
  // max := (*A)[i]
  (*A)[i] = (*A)[len(*A)-1]
  maxHeapify(*A, len(*A) - 1, i)
  return nil
}
