package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func add(v1 int, v2 int) (int, int) {
	res := v1 + v2
	c := 0

	if res > 9 {
		c = 1
	}

	return res % 10, c
}

func getNext(node *ListNode) *ListNode {
	if node != nil {
		return node.Next
	} else {
		return node
	}
}

func getValue(node *ListNode) int {
	if node == nil {
		return 0
	} else {
		return node.Val
	}
}

// SOLUTION
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{Val: 0, Next: nil}
	currentNode := res
	node1, node2 := l1, l2
	carry := 0

	for node1 != nil || node2 != nil {
		r, c := add(getValue(node1), carry+getValue(node2))
		carry = c

		currentNode.Val = r

		node1 = getNext(node1)
		node2 = getNext(node2)

		if node1 != nil || node2 != nil {
			currentNode.Next = &ListNode{Val: 0, Next: nil}
			currentNode = currentNode.Next
		}
	}

	if carry == 1 {
		currentNode.Next = &ListNode{Val: 1, Next: nil}
	}

	return res
}
