package mysort

import(
    "testing"
)

func TestStackEmpty(t *testing.T) {
    s := NewStack(10)
    if s.StackEmpty(){
        return
    }
    t.Errorf("StackEmpty Error")
}

func TestPush(t *testing.T) {
    s := NewStack(10)
    s.Push(1)
    if s.S[0] != 1 {
        t.Errorf("StackPush error and s.S[0] is %d", s.S[0])
    }
}

func TestPop(t *testing.T) {
    s := NewStack(10)
    err := s.Push(1)

    if err != nil {
        t.Errorf("StackPush error")
    }
    
    v, err := s.Pop()

    if err != nil || v != 1 {
        t.Errorf("StackPop error")
    }
}

func TestStackEmptyL(t *testing.T) {
    s := NewStackL()
    if s.StackEmpty(){
        return
    }
    t.Errorf("StackEmptyL Error")
}

func TestPushL(t *testing.T) {
    s := NewStackL()
    s.Push(1)
    if s.S.val != 1 {
        t.Errorf("StackPushL error and s.S[0] is %d", s.S.val)
    }
}

func TestPopL(t *testing.T) {
    s := NewStackL()
    err := s.Push(1)

    if err != nil {
        t.Errorf("StackPushL error")
    }
    
    v, err := s.Pop()

    if err != nil || v != 1 {
        t.Errorf("StackPopL error")
    }
}
