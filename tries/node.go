package tries

type Node struct {
	term     bool // true if from path to here a word in the trie
	children map[rune]*Node
}

// New creates an empty trie.
func New() *Node {
	node := Node{false, make(map[rune]*Node)}
	return &node
}
