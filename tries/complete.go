package tries

// AllWords returns a slice of all words in a trie.
func (n *Node) AllWords() []string {
	var words []string
	if n.term {
		words = []string{""}
	} else {
		words = make([]string, 0)
	}
	for r, node := range n.children {
		for _, word := range node.AllWords() {
			words = append(words, string(r)+word)
		}
	}
	return words
}

// Complete returns a slice of all words in a trie with a given prefix.
func (n *Node) Complete(pref string) []string {
	p := prefRoot(pref, n)
	if p == nil {
		return make([]string, 0)
	}
	words := p.AllWords()
	for i, s := range words {
		words[i] = pref + s
	}
	return words
}
