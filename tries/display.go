package tries

import "strings"

func (n *Node) String() string {
	return strings.Join(n.AllWords(), "\n")
}
