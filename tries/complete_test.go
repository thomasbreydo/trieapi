package tries

import (
	"sort"
	"testing"
)

func checkEq(desired []string, actual []string, t *testing.T) {
	if len(desired) != len(actual) {
		t.Errorf("desired: %q\nactual: %q", desired, actual)
		return
	}
	sort.Strings(desired)
	sort.Strings(actual)
	for i, s := range desired {
		if actual[i] != s {
			t.Errorf("desired: %q\nactual: %q", desired, actual)
			return
		}
	}
}

func TestComplete(t *testing.T) {
	trie := New()
	checkEq([]string{}, trie.Complete("a"), t)
	checkEq([]string{}, trie.Complete("Ӝ"), t)
	checkEq([]string{}, trie.Complete(""), t)
	trie.Add("test")
	checkEq([]string{}, trie.Complete("a"), t)
	checkEq([]string{}, trie.Complete("Ӝ"), t)
	checkEq([]string{"test"}, trie.Complete(""), t)
	checkEq([]string{"test"}, trie.Complete("t"), t)
	checkEq([]string{"test"}, trie.Complete("te"), t)
	checkEq([]string{"test"}, trie.Complete("tes"), t)
	checkEq([]string{"test"}, trie.Complete("test"), t)
	checkEq([]string{}, trie.Complete("tests"), t)
	checkEq([]string{}, trie.Complete("ter"), t)
}
