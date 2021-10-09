package main

// Comment it before submitting
type ListNode struct {
	data string
	next *ListNode
}

func Solution(head *ListNode, elem string) int {
	var i int
	for {
		if head.data == elem {
			return i
		}
		if head.next == nil {
			return -1
		}
		head = head.next
		i++
	}
}

// func test() {
// 	node3 := ListNode{"node3", nil}
// 	node2 := ListNode{"node2", &node3}
// 	node1 := ListNode{"node1", &node2}
// 	node0 := ListNode{"node0", &node1}
// 	/*idx :=*/ fmt.Println(Solution(&node0, "node1"))
// 	// result is : idx == 2
// }

// func main() {
// 	test()
// }
