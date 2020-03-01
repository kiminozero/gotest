package mysort

import(
    "testing"
    //"fmt"
)

func TestEnQueue(t *testing.T) {
    q := NewQueue(10)
    q.EnQueue(2)
    if q.Q[0] != 2 {
        t.Errorf("EnQueue error")
    }
}

func TestDeQueue(t *testing.T) {
    q := NewQueue(10)
    q.EnQueue(2)
    if q.Q[0] != 2 {
        t.Errorf("EnQueue error")
    }
    v, err := q.DeQueue()
    if 2 != v || err != nil {
        t.Errorf("DeQueue error")
    }
}
