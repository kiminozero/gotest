package mysort

import(
//    "fmt"
)
const INT_MAX = int(^uint(0)>>1)

const INT_MIN = ^INT_MAX

func merge(A []int, p, q, r int){
    n1 := q-p+1
    n2 := r-q
    
    L := make([]int, n1+1)
    R := make([]int, n2+1)
    
    for i := 0 ; i < n1; i++ {
        L[i] = A[p+i]
    }
    for j := 0 ; j < n2; j++ {
        R[j] = A[q+j+1]
    }
    L[n1] = INT_MAX
    R[n2] = INT_MAX
    i, j := 0, 0
    for k := p; k <= r; k++ {
        if L[i] < R[j] {
            A[k]= L[i]
            i++
        }else{
            A[k] = R[j]
            j++
        }
    }
}

func merge_noflag(A []int, p, q, r int){
    n1 := q-p+1
    n2 := r-q
    
    L := make([]int, n1)
    R := make([]int, n2)
    
    for i := 0 ; i < n1; i++ {
        L[i] = A[p+i]
    }
    for j := 0 ; j < n2; j++ {
        R[j] = A[q+j+1]
    }
    i, j, k := 0, 0, 0
    for k = p; k <= r; k++ {
        if i < n1 && j < n2{   //not the end of n1 and n2
                if L[i] < R[j] {
                    A[k]= L[i]
                    i++
                }else{
                    A[k] = R[j]
                    j++
                }
        }else{
            break
        }
    }
    if i < n1 {
        for i< n1{
            A[k] = L[i]
            k++ 
            i++
        }
    }else{
        for j< n2{
            A[k] = R[j]
            k++ 
            j++
        }
    }
}
func MergeSort(A []int, p, r int){
    // fmt.Println("Merge", p, (p+r)/2, r)
    if p < r {
        q := (p + r)/2
        MergeSort(A, p, q)
        MergeSort(A, q+1, r)
        // merge(A, p, q, r)
        merge_noflag(A, p, q, r)
    }
}
