package main

func PreOrderSearch(head BinaryNode) []int {
	return *walkPre(&head, &[]int{})
}

func walkPre(curr *BinaryNode, path *[]int) *[]int {
	if curr == nil {
		return path
	}

	*path = append(*path, curr.value)

	walkPre(curr.left, path)
	walkPre(curr.right, path)

	return path
}

func InOrderSearch(head BinaryNode) []int {
	return *walkIn(&head, &[]int{})
}

func walkIn(curr *BinaryNode, path *[]int) *[]int {
	if curr == nil {
		return path
	}

	walkIn(curr.left, path)
	*path = append(*path, curr.value)
	walkIn(curr.right, path)

	return path
}

func PostOrderSearch(head BinaryNode) []int {
	return *walkPost(&head, &[]int{})
}

func walkPost(curr *BinaryNode, path *[]int) *[]int {
	if curr == nil {
		return path
	}

	walkPost(curr.left, path)
	walkPost(curr.right, path)
	*path = append(*path, curr.value)

	return path
}

func Compare(a, b *BinaryNode) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	if a.value != b.value {
		return false
	}
	return Compare(a.left, b.left) && Compare(a.right, b.right)
}
