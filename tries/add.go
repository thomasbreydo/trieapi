package tries

// Add adds a keyword to the to a trie and returns true if the trie was modified
func (root *Node) Add(word string) bool {
	cur := root
	var next *Node
	var found bool
	modified := false
	for _, r := range word {
		next, found = cur.children[r]
		if !found {
			next = New()
			cur.children[r] = next
			modified = true
		}
		cur = next
	}
	if !cur.terminal {
		modified = true
	}
	cur.terminal = true
	return modified
}
