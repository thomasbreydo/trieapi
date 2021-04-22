package tries

// Delete removes a word from a trie and returns true if the trie was modified
func (n *Node) Delete(word string) bool {
	cur := n
	var found bool
	path := []*Node{n}
	for _, r := range word {
		cur, found = cur.children[r]
		if !found {
			return false // word not present
		}
		path = append(path, cur)
	}
	if !cur.term {
		return false
	}
	rr := []rune(word)
	for i := len(path) - 2; i >= 0; i-- {
		if path[i+1].term || len(path[i+1].children) > 0 {
			break
		}
		delete(path[i].children, rr[i])
	}
	cur.term = false
	return true
}
