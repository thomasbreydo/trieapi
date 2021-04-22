package tries

// rootOfPrefix finds the root of the subtrie of words that start with a prefix.
func rootOfPrefix(prefix string, root *Node) *Node {
	cur := root
	var found bool
	for _, r := range prefix {
		cur, found = cur.children[r]
		if !found {
			return nil
		}
	}
	return cur
}

// Search checks if a string is in a trie.
func (root *Node) Search(word string) bool {
	prefixRoot := rootOfPrefix(word, root)
	if prefixRoot == nil {
		return false
	}
	return prefixRoot.terminal
}
