package tries

import "testing"

func TestClear(t *testing.T) {
	trie := New()
	trie.Clear()
	trie.Add("hi")
	checkEq([]string{"hi"}, allWords(trie), t)
	trie.Clear()
	checkEq([]string{}, allWords(trie), t)
	trie.Clear()
	checkEq([]string{}, allWords(trie), t)
	trie.Add("hi")
	trie.Add("h")
	trie.Add("")
	trie.Add("*")
	trie.Add("Ш")
	checkEq([]string{"hi", "h", "", "*", "Ш"}, allWords(trie), t)
	trie.Clear()
	checkEq([]string{}, allWords(trie), t)
}
