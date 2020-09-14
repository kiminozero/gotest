package main

import (
    "fmt"
)

func buildNext(P string) []int {
    m := len(P)
    j := 0
    next := make([]int, m)
    next[0] = -1
    t := next[0]

    for j < m-1 {
        if 0 > t || P[j] == P[t] {
            j++
            t++
            next[j] = t
        } else {
            t = next[t]
        }
    }
    return next
}

func KMPSearch(S, P string) ([]int, bool) {
    i, n := 0, len(S)
    j, m := 0, len(P)
    next := buildNext(P)
    res := []int{}
    for i < n {
        if j == m-1 && S[i]==P[j] {
            // 匹配成功1次
            fmt.Println(" Found P at ", i-m+1)
            res = append(res, i-j)
            j = next[j]
        }
        if S[i] == P[j] {
            i++
            j++
        } else {
            j = next[j]
            if j == -1 {
                i++
                j++
            }
        }
    }
    if len(res) == 0 {
        return []int{-1}, false
    }
    return res, true
}

func main() {
    P1 := "abababac"
    S := "abcabcabcdabcd"
    P := "abcd"
    n := buildNext(P)
    fmt.Println(n)
    fmt.Println(buildNext(P1))
    fmt.Println(KMPSearch(S, P))
}
