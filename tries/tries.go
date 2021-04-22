package trie

type Node struct {
	terminal bool // true if from path from root to here is a word in the trie
	children map[rune]*Node
}

// New creates a trie root node.
func New(terminal bool) Node {
	return Node{terminal, make(map[rune]*Node)}
}

// Contains checks if a string is in a trie
func Contains(word string, root *Node) bool {
	var found bool
	cur := root
	for _, c := range word {
		cur, found = (*cur).children[c]
		if !found {
			return false
		}
	}
	return true
}

// Insert adds a string to the to a trie
func Insert(word string, root *Node) {
	var found bool
	cur := root
	for _, c := range word {
		cur, found = (*cur).children[c]
		if !found {
			node := New(false)
			(*cur).children[c] = &node
			cur = &node
		}
	}
	cur.terminal = true
}

// Delete removes a string from a trie, if present
func Delete(word string, root *Node) {
	var found bool
	cur := root
	for _, c := range word {
		cur, found = (*cur).children[c]
		if !found {
			return // string not present
		}
	}
	cur.terminal = false
}

// Suggest returns a slice of all words in a trie with a given prefix
func Suggest(prefix string)
