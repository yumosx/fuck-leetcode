/*
相交链表, 主要使用路径的计算
遍历a 链表, 遍历的路径是a -> b -> c
遍历b 链表, 遍历的路径是n -> m -> k

当a 遍历到空的时候, 我们让a 从b的起点开始遍历
当b 遍历到空的时候, 我们让b 从a 的起点开始遍历
*/
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	p := headA
	q := headB

	for p != q {
		if p != nil {
			p = p.Next
		} else {
			p = headB
		}

		if q != nil {
			q = q.Next
		} else {
			q = headA
		}
	}

	return p
}