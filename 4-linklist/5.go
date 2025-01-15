func detectCycle(head *ListNode) *ListNode {
	d := head
	d1 := head

	for d1 != nil && d1.Next != nil {
		d1 = d1.Next.Next
		d = d.Next

		if d == d1 {
			i := head
			j := d

			for i != j {
				i = i.Next
				j = j.Next
			}

			return i
		}
	}

	return head
}