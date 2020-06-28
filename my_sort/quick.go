package mysort

import(
    //"fmt"
    "math/rand"
)

func QuickSort(A []int, p, r int){
    if p<r{
        q := Partition(A, p, r)
        QuickSort(A, p, q-1)
        QuickSort(A, q+1, r)
    }
 
}


func Partition(A []int, p, r int) int{
    x := A[r]
    i := p - 1
    for j := p; j < r; j++ {
        if A[j] <= x {
            i++
            A[i], A[j] = A[j], A[i]
        }
    }
    A[i+1], A[r] = A[r], A[i+1]
    return i + 1
}

func RandomizedPartition(A []int, p, r int) int{
    /*if p >= r {
        return Partition(A, p, r)
    } */
    i := p + rand.Intn(r-p)
    A[r], A[i] = A[i], A[r]
    return Partition(A, p, r)
}

func RandomizedQuickSort(A []int, p, r int) {
    if p < r {
        q := RandomizedPartition(A, p, r)
        RandomizedQuickSort(A, p, q-1)
        RandomizedQuickSort(A, q+1, r)
    }

}

func QuickSortS(A []int, p, r int) {
    for p<r {
        // Partition and sort left subarray
        q := Partition(A, p, r)
        QuickSortS(A, p, q-1)
        p = q+1
    }
}

func RandomizedSelect(A []int, p, r, i int) int{
    if p==r{
        return A[p]
    }
    q := RandomizedPartition(A, p ,r)
    k := q-p+1
    if i==k {
        return A[q]
    }else if i<k {
        return RandomizedSelect(A, p, q-1, i)
    }else{
        return RandomizedSelect(A, q+1, r, i-k)
    }
}

func RandomizedSelectI(A []int, p, r, i int) int{

    for p != r {
        q := RandomizedPartition(A, p ,r)
        k := q-p+1
        if i<k {
            r = q - 1
        }else if i > k{
            p = q + 1
            i = i-k
        }else{
            return A[q]
        }
    }

    return A[p]
}



