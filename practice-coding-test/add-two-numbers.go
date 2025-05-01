package main

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	tmp := result
	carry := 0

	for l1 != nil || l2 != nil || carry != 0 {
		first, second := 0, 0

		if l1 != nil {
			first = l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			second = l2.Val
			l2 = l2.Next
		}

		sum := first + second + carry
		carry = sum / 10

		node := &ListNode{Val: sum % 10}

		tmp.Next = node
		tmp = tmp.Next
	}

	return result.Next
}
