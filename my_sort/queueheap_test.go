package mysort

import(
  "testing"
)

// Minheap can implement queue using counter time.now() as the key
// k first node make minheap S, Minimum(S) if min next != nil than S[0]= S[0]next, minheapify(S) else than Extract-Min(S)



func TestHeapMaximum(t *testing.T){
    nums := []int{1,0,3,4,2,5,8,9,7,6}
    buildMaxHeap(nums)
    max := HeapMaximum(nums)
    if max != 9{
        t.Errorf("HeapMaximum wrong %d", max)
    }
}

func TestHeapExtractMax(t *testing.T){
    nums := []int{1,0,3,4,2,5,8,9,7,6}
    buildMaxHeap(nums)
    max := HeapMaximum(nums)
    if max != 9{
        t.Errorf("HeapMaximum wrong %d", max)
    }
    max, _ = HeapExtractMax(nums)
    if max != 9{
        t.Errorf("HeapMaximum wrong %d", max)
    }
    max = HeapMaximum(nums)
    if max != 8{
        t.Errorf("HeapMaximum wrong %d", max)
    }
    
}

func TestHeapIncreaseKey(t *testing.T){
    nums := []int{1,0,3,4,2,5,8,9,7,6}
    buildMaxHeap(nums)
    max := HeapMaximum(nums)
    if max != 9{
        t.Errorf("HeapMaximum wrong %d", max)
    }
    err := HeapIncreaseKey(nums, 9, 10)
    if err != nil {
        t.Errorf("%s", err)
    }
    max = HeapMaximum(nums)
    if max != 10 {
        t.Errorf("HeapMaximum not 10 is %d", max)
    }
}

func TestHeapInsert(t *testing.T){
    nums := []int{1,0,3,4,2,5,8,9,7,6}
    buildMaxHeap(nums)
    max := HeapMaximum(nums)
    if max != 9{
        t.Errorf("HeapMaximum wrong %d", max)
    }
    MaxHeapInsert(&nums, 10)
    
    max = HeapMaximum(nums[:])
    if max != 10 {
        t.Errorf("HeapMaximum not 10 is %d", max)
    }
}
