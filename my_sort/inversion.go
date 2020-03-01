package mysort
import(
//    "fmt"
    )

func merge_noflag1(A []int, p, q, r int, c *int){
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
                    *c+=n1-i
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
func Inversion(A []int, p, r int, c *int){
    // c  count inversion counts.
    // fmt.Println("Merge", p, (p+r)/2, r)
    
    if p < r {
        q := (p + r)/2
        Inversion(A, p, q, c)
        Inversion(A, q+1, r, c)
        // merge(A, p, q, r)
        merge_noflag1(A, p, q, r, c)
    }
  //  fmt.Println(c)
}
