package main

import "fmt"

// Comment it before submitting
type ListNode struct {
	data string
	next *ListNode
}

func Solution(head *ListNode, idx int) *ListNode {
	if idx == 0 {
		if head.next != nil {
			return head.next
		}
		return head
	}

	var previous, next *ListNode
	var err error
	if previous, err = getNodeByIndex(head, idx-1); err != nil {
		return head
	}
	if next, err = getNodeByIndex(head, idx+1); err == nil {
		previous.next = next
	} else {
		previous.next = nil
	}

	return head
}

func getNodeByIndex(head *ListNode, idx int) (*ListNode, error) {
	if idx > 0 && head.next != nil {
		return getNodeByIndex(head.next, idx-1)
	}

	if idx > 0 && head.next == nil {
		return nil, fmt.Errorf("index out")
	}
	return head, nil
}

// func test() {
// 	node3 := ListNode{"node3", nil}
// 	node2 := ListNode{"node2", &node3}
// 	node1 := ListNode{"node1", &node2}
// 	node0 := ListNode{"node0", &node1}
// 	newHead := Solution(&node0, 1)

// 	Print(newHead)

// 	// result is : node0 -> node2 -> node3
// }

// func Print(head *ListNode) {
// 	fmt.Println(head.data)
// 	if head.next != nil {
// 		Print(head.next)
// 	}
// }

// func main() {
// 	test()
// }
