package tries

import "testing"

func TestNew(t *testing.T) {
	trie := New()
	if trie.term {
		t.Error("New().terminal is (true) and not (false)")
	}
	if len(trie.children) != 0 {
		t.Errorf("New().children has length (%d) and not 0", len(trie.children))
	}
}
