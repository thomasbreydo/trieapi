package tries

func (root *Node) Clear() {
	root.term = false
	root.children = make(map[rune]*Node)
}
