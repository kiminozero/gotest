package mysort

import(
    "fmt"
    )

func maxHeapify(A []int, size, i int){
  l := 2*i + 1  // LEFT(i)
  r := 2*i + 2    // RIGHT(i)
  largest := i
  if l< size && A[l] > A[largest]{
    largest = l
  }

  if r< size && A[r] > A[largest]{
    largest = r
  }

  if largest != i {
    A[i], A[largest] = A[largest], A[i]
    maxHeapify(A, size, largest)
  }
}

func buildMaxHeap(A []int){
  heapsize := len(A)
  for i := heapsize/2-1; i >= 0; i-- {
    // fmt.Println("i",i, "A", A)
    maxHeapify(A, heapsize, i)
    // fmt.Println("i",i, "A", A)
  }
  fmt.Println(A)
}

func buildMinHeap(A []int){
  heapsize := len(A)
  for i := heapsize/2-1; i >= 0; i-- {
    // fmt.Println("i",i, "A", A)
    minHeapify(A, heapsize, i)
    // fmt.Println("i",i, "A", A)
  }
  fmt.Println(A)
}

func minHeapify(A []int, size, i int){
  l := 2*i + 1  // LEFT(i)
  r := 2*i + 2    // RIGHT(i)
  smallest := i
  if l< size && A[l] < A[smallest]{
    smallest = l
  }

  if r< size && A[r] < A[smallest]{
    smallest = r
  }

  if smallest != i {
    A[i], A[smallest] = A[smallest], A[i]
    maxHeapify(A, size, smallest)
  }
}

func HeapSort(A []int){
  buildMaxHeap(A)
  for i := len(A)-1; i > 0; i-- {
    A[0], A[i] = A[i], A[0]
    maxHeapify(A, i, 0)
  }
  fmt.Println(A)
}

func HeapSortReverse(A []int){
  buildMinHeap(A)
  for i := len(A)-1; i > 0; i-- {
    A[0], A[i] = A[i], A[0]
    minHeapify(A, i, 0)
  }
  fmt.Println(A)
}
