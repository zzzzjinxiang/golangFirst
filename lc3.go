package main

import "net/http"

type ListNode struct {
	Val  int
	Next *ListNode
}

var (
	ltmp *ListNode
	cnt  int = 0
)

func main() {

	http.ListenAndServe("127.0.0.1", nil)
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	var head = &ListNode{
		Val:  -1,
		Next: nil,
	};
	pHead := head
	carry := 0;
	for l1 != nil || l2 != nil {

		tmp := 0;
		if l1 != nil && l2 != nil {
			tmp = l1.Val + l2.Val;
			l1 = l1.Next;
			l2 = l2.Next;
		} else if l1 != nil {
			tmp = l1.Val;
			l1 = l1.Next;
		} else {
			tmp = l2.Val;
			l2 = l2.Next;
		}

		head.Next = &ListNode{
			Val:  (tmp+carry)%10,
			Next: nil,
		};

		carry = (carry+tmp) / 10;

		head = head.Next;

	}
	if (carry != 0) {
		head.Next = &ListNode{
			Val:  carry,
			Next: nil,
		};
		head = head.Next;
	}

	return pHead.Next;
}
