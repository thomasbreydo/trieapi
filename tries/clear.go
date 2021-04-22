package tries

func (n *Node) Clear() {
	n.term = false
	n.children = make(map[rune]*Node)
}
