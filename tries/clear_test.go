package tries

import (
	"testing"
)

func TestClear(t *testing.T) {
	trie := New()
	trie.Clear()
	trie.Add("hi")
	checkEq([]string{"hi"}, trie.AllWords(), t)
	trie.Clear()
	checkEq([]string{}, trie.AllWords(), t)
	trie.Clear()
	checkEq([]string{}, trie.AllWords(), t)
	trie.Add("hi")
	trie.Add("h")
	trie.Add("")
	trie.Add("*")
	trie.Add("Ш")
	checkEq([]string{"hi", "h", "", "*", "Ш"}, trie.AllWords(), t)
	trie.Clear()
	checkEq([]string{}, trie.AllWords(), t)
}
