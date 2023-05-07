package main

func Bfs(head BinaryNode, needle int) bool {
	var queue DoublyLinkedList[BinaryNode]
	queue.Append(head)

	for queue.length != 0 {
		curr := queue.Pop()

		if curr.value == needle {
			return true
		}

		if curr.left != nil {
			queue.Append(*curr.left)
		}
		if curr.right != nil {
			queue.Append(*curr.right)
		}
	}

	return false
}
