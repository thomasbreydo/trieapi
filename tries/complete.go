package tries

// allWords returns a slice of all words in a trie.
func allWords(root *Node) []string {
	var words []string
	if root.terminal {
		words = []string{""}
	} else {
		words = make([]string, 0)
	}
	for r, node := range root.children {
		for _, word := range allWords(node) {
			words = append(words, string(r)+word)
		}
	}
	return words
}

// Complete returns a slice of all words in a trie with a given prefix.
func (root *Node) Complete(prefix string) []string {
	prefixRoot := rootOfPrefix(prefix, root)
	if prefixRoot == nil {
		return make([]string, 0)
	}
	words := allWords(prefixRoot)
	for i, s := range words {
		words[i] = prefix + s
	}
	return words
}
