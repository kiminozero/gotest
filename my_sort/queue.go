package mysort

import(
    "fmt"
)

type Queue interface {
    EnQueue()
    DeQueue()
}

// Dueue use one []int
type Dueue interface {
    LEnQueue()
    LDeQueue()
    REnQueue()
    RDeQueue()
}


type queue struct {
    Q        []int
    headQ    int
    tailQ    int
}

func NewQueue(size int) *queue {
    q := queue{make([]int, size), 0,0}
    return &q
}

func (q *queue)EnQueue(x int) error {

    if q.headQ == q.tailQ + 1 {
        return fmt.Errorf("overflow")
    }
    q.Q[q.tailQ] = x
    if q.tailQ == len(q.Q)-1 {
        q.tailQ = 0
    }else{
        q.tailQ++
    }
    return nil
}

func (q *queue)DeQueue() (int, error) {

    if q.headQ == q.tailQ {
        return 0 , fmt.Errorf("underflow")
    }
    x := q.Q[q.headQ]
    if q.headQ == len(q.Q)-1 {
        q.headQ = 1
    }else{
        q.headQ++
    }
    return x, nil
}

