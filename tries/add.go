package tries

// Add adds a word to the to a trie and returns true if the trie was modified
func (n *Node) Add(w string) bool {
	cur := n
	var next *Node
	var found bool
	mod := false
	for _, r := range w {
		next, found = cur.children[r]
		if !found {
			next = New()
			cur.children[r] = next
			mod = true
		}
		cur = next
	}
	if !cur.term {
		mod = true
	}
	cur.term = true
	return mod
}
