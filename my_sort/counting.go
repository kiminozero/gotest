package mysort

import(
    //"fmt"
)

func CountingSort(A, B []int, k int) {
    C := make([]int, k)
    // counting every num count
    for _, j := range(A){
        C[j] = C[j] + 1
    }
    //fmt.Println(C)
    // counting <= numi's count
    for i := 1; i < k; i++ {
        C[i] = C[i] + C[i-1]
    }
    // fmt.Println(C)

    for j := len(A)-1; j>=0; j-- {
        C[A[j]] = C[A[j]] - 1
        B[C[A[j]]] = A[j]
        // fmt.Println(A[j], B, C)
    }
}



func CountingRange(A []int, k ,a, b int) int{
    C := make([]int, k)
    // counting every num count
    for _, j := range(A){
        C[j] = C[j] + 1
    }
    //fmt.Println(C)
    // counting <= numi's count
    for i := 1; i < k; i++ {
        C[i] = C[i] + C[i-1]
    }
    // fmt.Println(C)

    return C[b]-C[a-1] 
}


