package easy

type ListNode struct {
	Val  int
	Next *ListNode
}

func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 != nil {
		return l2
	}
	if l1 != nil && l2 == nil {
		return l1
	}
	if l1 == nil && l2 == nil {
		return nil
	}
	newNode := new(ListNode)
	if l1.Val >= l2.Val {
		newNode.Val = l2.Val
		l2 = l2.Next
		newNode.Next = MergeTwoLists(l1, l2)
	} else {
		newNode.Val = l1.Val
		l1 = l1.Next
		newNode.Next = MergeTwoLists(l1, l2)
	}
	return newNode
}
