package tries

// Clear removes all words from this trie and returns true if the trie was modified.
func (n *Node) Clear() bool {
	mod := n.term || len(n.children) > 0
	n.term = false
	n.children = make(map[rune]*Node)
	return mod
}
