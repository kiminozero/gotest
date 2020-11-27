package main
// Treap
import (
    "math/rand"
    "fmt"
)

type node struct {
    ch       [2]*node
    priority int
    key      int
    dupCnt   int
    sz       int
}

func (o *node) cmp(b int) int {
    switch {
    case b < o.key:
        return 0
    case b > o.key:
        return 1
    default:
        return -1
    }
}

func (o *node) size() int {
    if o != nil {
        return o.sz
    }
    return 0
}

func (o *node) maintain() {
    o.sz = o.dupCnt + o.ch[0].size() + o.ch[1].size()
}

func (o *node) rotate(d int) *node {
    x := o.ch[d^1]
    o.ch[d^1] = x.ch[d]
    x.ch[d] = o
    o.maintain()
    x.maintain()
    return x
}

type treap struct {
    root *node
}

func (t *treap) _insert(o *node, key int) *node {
    if o == nil {
        return &node{priority: rand.Int(), key: key, dupCnt: 1, sz: 1}
    }
    if d := o.cmp(key); d >= 0 {
        o.ch[d] = t._insert(o.ch[d], key)
        if o.ch[d].priority > o.priority {
            o = o.rotate(d ^ 1)
        }
    } else {
        o.dupCnt++
    }
    o.maintain()
    return o
}

func (t *treap) insert(key int) {
    t.root = t._insert(t.root, key)
}

// equal=false: 小于 key 的元素个数
// equal=true: 小于或等于 key 的元素个数
func (t *treap) rank(key int, equal bool) (cnt int) {
    for o := t.root; o != nil; {
        switch c := o.cmp(key); {
        case c == 0:
            o = o.ch[0]
        case c > 0:
            cnt += o.dupCnt + o.ch[0].size()
            o = o.ch[1]
        default:
            cnt += o.ch[0].size()
            if equal {
                cnt += o.dupCnt
            }
            return
        }
    }
    return
}

func countRangeSum(nums []int, lower, upper int) (cnt int) {
    preSum := make([]int, len(nums)+1)
    for i, v := range nums {
        preSum[i+1] = preSum[i] + v
    }

    t := &treap{}
    for _, sum := range preSum {
        left, right := sum-upper, sum-lower
        cnt += t.rank(right, true) - t.rank(left, false)
        t.insert(sum)
    }
    return
}

// 统计左侧小于当前元素的个数
func main() {
    nums := []int{1,2,6,3,7,4,7,9,2,0,8}
    res := make([]int, len(nums))
    t := &treap{}
    for i := range nums {
        rank := t.rank(nums[i],false) 
        res[i] = rank
        t.insert(nums[i])
    }

    fmt.Println(nums)
    fmt.Println(res)

}
