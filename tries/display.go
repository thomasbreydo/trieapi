package tries

import "strings"

func (root *Node) String() string {
	return strings.Join(allWords(root), "\n")
}
