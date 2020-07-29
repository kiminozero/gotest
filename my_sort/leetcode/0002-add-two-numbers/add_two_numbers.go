package leetcode


//  Definition for singly-linked list.

 type ListNode struct {
     Val int
     Next *ListNode
 }
 
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    carry := 0
    head := &ListNode{}
    cur := head
    for l1 != nil ||l2 != nil || carry != 0 {
        a, b := 0, 0
        if l1 !=nil {
            a = l1.Val
            l1 = l1.Next
        }
        if l2 != nil {
            b = l2.Val
            l2 =l2.Next
        }
        cur.Next = &ListNode{(carry+ a + b)%10, nil}
        carry = (carry + a + b)/10
        cur = cur.Next
    }

    return head.Next
}
