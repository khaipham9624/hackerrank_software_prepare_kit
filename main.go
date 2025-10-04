package main

import (
	"fmt"
)

func printAllNode(head *SinglyLinkedListNode) {
	if head == nil {
		return
	}
	tmpNode := head
	for {
		fmt.Print(tmpNode.data, " ")
		if tmpNode.next != nil {
			tmpNode = tmpNode.next
		} else {
			break
		}
	}
}

func deleteDuplicates(head *SinglyLinkedListNode) *SinglyLinkedListNode {
	// Write your code here
	if head == nil {
		return head
	}
	tmpNode := head
	for {
		if tmpNode.next.data == tmpNode.data {
			tmpNode = tmpNode.next

		}
	}
	return head
}

func main() {
	head := nil
	res := deleteDuplicates(head)
}
