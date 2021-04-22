package tries

import "strings"

func (root *Node) String() string {
	// todo tree viz
	return strings.Join(allWords(root), "\n")
}
