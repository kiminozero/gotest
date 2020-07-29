package leetcode
import (
	"fmt"
	"testing"
)
func getListFromSlice(nums []int) *ListNode  {
	if len(nums) == 0{
		return nil
	}
	head := &ListNode{nums[0], nil}
	cur := head
	for i := range nums {
		if i == 0 {
			continue
		}
		cur.Next = &ListNode{nums[i], nil}
		cur = cur.Next
	}
	return head
}

func travalList(head *ListNode) []int  {

	res := []int{}
	if head == nil {
		return []int{}
	}
    for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}
	return res
}
type testData struct {
    a, b, c   []int
}
func TestAddTwoNumbers(t *testing.T)  {
    tests := []testData{
            testData{[]int{2,4,3}, []int{5,6,4}, []int{7,0,8}},
    }
    for _, testcase := range tests {
	    a, b, c := testcase.a, testcase.b, testcase.c
	    res := addTwoNumbers(getListFromSlice(a), getListFromSlice(b))
        for i, v := range travalList(res) {
            // fmt.Println(v, c[i])
            if c[i] != v {
                t.Errorf("addTwoNumbers failed at case %d" , i+1)
                fmt.Println(a, "\n", b, "\n", c)
                break
            }
        }

    }

}
