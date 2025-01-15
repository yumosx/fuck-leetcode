func midddleNode(head *ListNode) *ListNode {
	d := head
	d1 := head

	for d1 != nil && d1.Next != nil {
		d = d.Next
		d1 = d1.Next.Next
	}

	//返回链表的中间节点
	return d
}

func reverseList(head *ListNode) *ListNode {
	d := head
	d1 := head.Next

	for d1 != nil {
		tmp := d1.Next

		d1.Next = d
		d = d1
		d1 = tmp
	}

	head.Next = nil
	head = d
	return head
}

func isPalindrome(head *ListNode) bool {
	mid := midddleNode(head)
	l := reverseList(mid)

	for l != nil {
		if l.Val == head.Val {
			l = l.Next
			head = head.Next
		} else {
			return false
		}
	}

	return true
}