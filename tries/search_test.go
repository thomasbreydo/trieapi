package tries

import "testing"

func TestSearch(t *testing.T) {
	trie := New()
	if trie.Search("") {
		t.Error("trie.Search is (true) and not (false)")
	}
	if trie.Search("asfj") {
		t.Error("trie.Search is (true) and not (false)")
	}
	if trie.Search("Ӝ") {
		t.Error("trie.Search is (true) and not (false)")
	}
	trie.Add("Ӝ")
	if trie.Search("") {
		t.Error("trie.Search is (true) and not (false)")
	}
	if trie.Search("asfj") {
		t.Error("trie.Search is (true) and not (false)")
	}
	if !trie.Search("Ӝ") {
		t.Error("trie.Search (false) and not (true)")
	}
	if trie.Search("ӜӜ") {
		t.Error("trie.Search is (true) and not (false)")
	}
	trie.Add("")
	if !trie.Search("") {
		t.Error("trie.Search (false) and not (true)")
	}
	if trie.Search("asfj") {
		t.Error("trie.Search is (true) and not (false)")
	}
	if !trie.Search("Ӝ") {
		t.Error("trie.Search (false) and not (true)")
	}
	if trie.Search("ӜӜ") {
		t.Error("trie.Search is (true) and not (false)")
	}
	trie.Add("asf")
	if !trie.Search("") {
		t.Error("trie.Search (false) and not (true)")
	}
	if trie.Search("asfj") {
		t.Error("trie.Search is (true) and not (false)")
	}
	if !trie.Search("Ӝ") {
		t.Error("trie.Search (false) and not (true)")
	}
	if trie.Search("ӜӜ") {
		t.Error("trie.Search is (true) and not (false)")
	}
	trie.Delete("nothere")
	if !trie.Search("") {
		t.Error("trie.Search (false) and not (true)")
	}
	if trie.Search("asfj") {
		t.Error("trie.Search is (true) and not (false)")
	}
	if !trie.Search("Ӝ") {
		t.Error("trie.Search (false) and not (true)")
	}
	if trie.Search("ӜӜ") {
		t.Error("trie.Search is (true) and not (false)")
	}
	trie.Delete("")
	if trie.Search("") {
		t.Error("trie.Search is (true) and not (false)")
	}
	if trie.Search("asfj") {
		t.Error("trie.Search is (true) and not (false)")
	}
	if !trie.Search("Ӝ") {
		t.Error("trie.Search (false) and not (true)")
	}
	if trie.Search("ӜӜ") {
		t.Error("trie.Search is (true) and not (false)")
	}
	trie.Add("ӜӜ")
	if trie.Search("") {
		t.Error("trie.Search is (true) and not (false)")
	}
	if trie.Search("asfj") {
		t.Error("trie.Search is (true) and not (false)")
	}
	if !trie.Search("Ӝ") {
		t.Error("trie.Search (false) and not (true)")
	}
	if !trie.Search("ӜӜ") {
		t.Error("trie.Search (false) and not (true)")
	}
}
