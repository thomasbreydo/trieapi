package tries

import (
	"testing"
)

func TestClear(t *testing.T) {
	trie := New()
	if trie.Clear() {
		t.Error("expected false")
	}
	trie.Add("hi")
	checkEq([]string{"hi"}, trie.AllWords(), t)
	if !trie.Clear() {
		t.Error("expected true")
	}
	checkEq([]string{}, trie.AllWords(), t)
	if trie.Clear() {
		t.Error("expected false")
	}
	checkEq([]string{}, trie.AllWords(), t)
	trie.Add("hi")
	trie.Add("h")
	trie.Add("")
	trie.Add("*")
	trie.Add("Ш")
	checkEq([]string{"hi", "h", "", "*", "Ш"}, trie.AllWords(), t)
	if !trie.Clear() {
		t.Error("expected true")
	}
	checkEq([]string{}, trie.AllWords(), t)
	if trie.Clear() {
		t.Error("expected false")
	}
}
