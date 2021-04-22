package tries

// prefRoot finds the root of the subtrie of words that start with a prefix.
func prefRoot(prefix string, root *Node) *Node {
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

// Search checks if a word is in a trie.
func (n *Node) Search(word string) bool {
	prefixRoot := prefRoot(word, n)
	if prefixRoot == nil {
		return false
	}
	return prefixRoot.term
}
