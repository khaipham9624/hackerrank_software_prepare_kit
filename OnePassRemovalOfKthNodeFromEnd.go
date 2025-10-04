package main

type SinglyLinkedListNode struct {
	data int32
	next *SinglyLinkedListNode
}

func removeKthNodeFromEnd(head *SinglyLinkedListNode, k int32) *SinglyLinkedListNode {
	// Write your code here
	if k < 0 {
		return head
	}
	stack := []*SinglyLinkedListNode{}

	var tmpNode *SinglyLinkedListNode = head
	for {
		stack = append([]*SinglyLinkedListNode{tmpNode}, stack...)
		if tmpNode.next == nil {
			break
		}
		tmpNode = tmpNode.next
	}
	if k >= int32(len(stack)) {
		return head
	}

	if len(stack) <= 1 {
		if k == 0 {
			return nil
		}
		if k == 1 {
			return head
		}
	}
	if k == 0 {
		stack[k+1].next = nil
	} else if k == int32(len(stack)-1) {
		return stack[k-1]
	} else {
		stack[k+1].next = stack[k-1]
	}
	return head
}
