package tries

import "testing"

func TestDelete(t *testing.T) {
	trie := New()
	if trie.Delete("") {
		t.Error("trie.Delete is (true) and not (false)")
	}
	if trie.Delete("Ӝ") {
		t.Error("trie.Delete is (true) and not (false)")
	}
	trie.Add("hi")
	if trie.Delete("") {
		t.Error("trie.Delete is (true) and not (false)")
	}
	if trie.Delete("h") {
		t.Error("trie.Delete is (true) and not (false)")
	}
	if !trie.Delete("hi") {
		t.Error("trie.Delete is (false) and not (true)")
	}
	if trie.Delete("hi") {
		t.Error("trie.Delete is (true) and not (false)")
	}
	trie.Add("hello")
	trie.Add("hellooӜ")
	trie.Add("hellooy")
	checkEq([]string{"hello", "hellooӜ", "hellooy"}, trie.AllWords(), t)
	if trie.Delete("") {
		t.Error("trie.Delete is (true) and not (false)")
	}
	checkEq([]string{"hello", "hellooӜ", "hellooy"}, trie.AllWords(), t)
	if trie.Delete("hi") {
		t.Error("trie.Delete is (true) and not (false)")
	}
	checkEq([]string{"hello", "hellooӜ", "hellooy"}, trie.AllWords(), t)
	if !trie.Delete("hello") {
		t.Error("trie.Delete is (false) and not (true)")
	}
	checkEq([]string{"hellooӜ", "hellooy"}, trie.AllWords(), t)
	trie.Add("fix")
	if !trie.Delete("hellooӜ") {
		t.Error("trie.Delete is (false) and not (true)")
	}
	checkEq([]string{"hellooy", "fix"}, trie.AllWords(), t)
	if trie.Delete("") {
		t.Error("trie.Delete is (true) and not (false)")
	}
	trie.Add("")
	checkEq([]string{"", "hellooy", "fix"}, trie.AllWords(), t)
	trie.Add("h")
	checkEq([]string{"", "h", "hellooy", "fix"}, trie.AllWords(), t)
	if !trie.Delete("h") {
		t.Error("trie.Delete is (false) and not (true)")
	}
	checkEq([]string{"", "hellooy", "fix"}, trie.AllWords(), t)
	trie.Add("h")
	checkEq([]string{"", "h", "hellooy", "fix"}, trie.AllWords(), t)
	if !trie.Delete("") {
		t.Error("trie.Delete is (false) and not (true)")
	}
	checkEq([]string{"h", "hellooy", "fix"}, trie.AllWords(), t)
	if !trie.Delete("h") {
		t.Error("trie.Delete is (false) and not (true)")
	}
	checkEq([]string{"hellooy", "fix"}, trie.AllWords(), t)
	if !trie.Delete("fix") {
		t.Error("trie.Delete is (false) and not (true)")
	}
	checkEq([]string{"hellooy"}, trie.AllWords(), t)
	if !trie.Delete("hellooy") {
		t.Error("trie.Delete is (false) and not (true)")
	}
	checkEq([]string{}, trie.AllWords(), t)
}
