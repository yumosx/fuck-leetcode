func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

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