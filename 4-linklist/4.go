func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	d := head
	d1 := head.Next

	for d != d1 {
		if d1 != nil && d1.Next != nil {
			d1 = d1.Next.Next
		} else {
			return false
		}

		d = d.Next
	}

	return true
}