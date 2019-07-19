/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func Abs(a, b int) int {
	if a > b {
		return a - b
	} else {
		return b - a
	}
}
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	s1, s2 := 0, 0
	for temp := l1; temp != nil; temp = temp.Next {
		s1++
	}
	for temp := l2; temp != nil; temp = temp.Next {
		s2++
	}
	var short, long *ListNode
	if s1 > s2 {
		long = l1
		short = l2
	} else {
		long = l2
		short = l1
	}
	var head, temp *ListNode
	for i := Abs(s1, s2); i != 0; i-- {
		temp = long.Next
		long.Next = head
		head = long
		long = temp
	}

	for long != nil {
		node := ListNode{Val: long.Val + short.Val, Next: head}
		head = &node

		long = long.Next
		short = short.Next
	}

	carry := 0
	for temp = head; temp != nil; temp = temp.Next {
		v := temp.Val + carry
		temp.Val = v % 10
		carry = v / 10
	}

	// reverse link list
	temp = head
	head = nil
	var tt *ListNode
	for temp != nil {
		tt = temp.Next
		temp.Next = head
		head = temp
		temp = tt
	}

	// handle carry
	if carry != 0 {
		nn := ListNode{Val: carry, Next: head}
		head = &nn
	}
	return head
}

