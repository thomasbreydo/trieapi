package tries

import "testing"

// TestAdd tests trie insertion, including unicode characters and duplicates
func TestAdd(t *testing.T) {
	trie := New()

	if !trie.Add("Ӝ∮∮") {
		t.Error("trie.Add is (false) and not (true)")
	}
	// trie contains: Ӝ∮∮
	var ok bool
	var cur *Node
	cur = trie
	if cur.term {
		t.Error("cur.terminal is (true) and not (false)")
	}
	if len(cur.children) != 1 {
		t.Errorf("cur.children has length %d and not 1", len(cur.children))
	}
	cur, ok = cur.children['Ӝ']
	if !ok {
		t.Error("cur.children doesn't contain Ӝ")
	}
	if cur.term {
		t.Error("cur.terminal is (true) and not (false)")
	}
	if len(cur.children) != 1 {
		t.Errorf("cur.children has length %d and not 1", len(cur.children))
	}
	cur, ok = cur.children['∮']
	if !ok {
		t.Error("cur.children doesn't contain ∮")
	}
	if cur.term {
		t.Error("cur.terminal is (true) and not (false)")
	}
	if len(cur.children) != 1 {
		t.Errorf("cur.children has length %d and not 1", len(cur.children))
	}
	cur, ok = cur.children['∮']
	if !ok {
		t.Error("cur.children doesn't contain ∮")
	}
	if !cur.term {
		t.Error("cur.terminal is (false) and not (true)")
	}
	if len(cur.children) != 0 {
		t.Errorf("cur.children has length %d and not 0", len(cur.children))
	}

	if !trie.Add("Ӝ∮") {
		t.Error("trie.Add is (false) and not (true)")
	}
	// trie contains: Ӝ∮∮, Ӝ∮
	cur = trie
	if cur.term {
		t.Error("cur.terminal is (true) and not (false)")
	}
	if len(cur.children) != 1 {
		t.Errorf("cur.children has length %d and not 1", len(cur.children))
	}
	cur, ok = cur.children['Ӝ']
	if !ok {
		t.Error("cur.children doesn't contain Ӝ")
	}
	if cur.term {
		t.Error("cur.terminal is (true) and not (false)")
	}
	if len(cur.children) != 1 {
		t.Errorf("cur.children has length %d and not 1", len(cur.children))
	}
	cur, ok = cur.children['∮']
	if !ok {
		t.Error("cur.children doesn't contain ∮")
	}
	if !cur.term {
		t.Error("cur.terminal is (false) and not (true)")
	}
	if len(cur.children) != 1 {
		t.Errorf("cur.children has length %d and not 1", len(cur.children))
	}
	cur, ok = cur.children['∮']
	if !ok {
		t.Error("cur.children doesn't contain ∮")
	}
	if !cur.term {
		t.Error("cur.terminal is (false) and not (true)")
	}
	if len(cur.children) != 0 {
		t.Errorf("cur.children has length %d and not 0", len(cur.children))
	}

	if !trie.Add("") {
		t.Error("trie.Add is (false) and not (true)")
	}
	// trie contains: Ӝ∮∮, <empty string>, Ӝ∮
	cur = trie
	if !cur.term {
		t.Error("cur.terminal is (false) and not (true)")
	}
	if len(cur.children) != 1 {
		t.Errorf("cur.children has length %d and not 1", len(cur.children))
	}
	cur, ok = cur.children['Ӝ']
	if !ok {
		t.Error("cur.children doesn't contain Ӝ")
	}
	if cur.term {
		t.Error("cur.terminal is (true) and not (false)")
	}
	if len(cur.children) != 1 {
		t.Errorf("cur.children has length %d and not 1", len(cur.children))
	}
	cur, ok = cur.children['∮']
	if !ok {
		t.Error("cur.children doesn't contain ∮")
	}
	if !cur.term {
		t.Error("cur.terminal is (false) and not (true)")
	}
	if len(cur.children) != 1 {
		t.Errorf("cur.children has length %d and not 1", len(cur.children))
	}
	cur, ok = cur.children['∮']
	if !ok {
		t.Error("cur.children doesn't contain ∮")
	}
	if !cur.term {
		t.Error("cur.terminal is (false) and not (true)")
	}
	if len(cur.children) != 0 {
		t.Errorf("cur.children has length %d and not 0", len(cur.children))
	}

	if trie.Add("") {
		t.Error("trie.Add is (true) and not (false)")
	}
	// trie contains: Ӝ∮∮, <empty string>, Ӝ∮
	cur = trie
	if !cur.term {
		t.Error("cur.terminal is (false) and not (true)")
	}
	if len(cur.children) != 1 {
		t.Errorf("cur.children has length %d and not 1", len(cur.children))
	}
	cur, ok = cur.children['Ӝ']
	if !ok {
		t.Error("cur.children doesn't contain Ӝ")
	}
	if cur.term {
		t.Error("cur.terminal is (true) and not (false)")
	}
	if len(cur.children) != 1 {
		t.Errorf("cur.children has length %d and not 1", len(cur.children))
	}
	cur, ok = cur.children['∮']
	if !ok {
		t.Error("cur.children doesn't contain ∮")
	}
	if !cur.term {
		t.Error("cur.terminal is (false) and not (true)")
	}
	if len(cur.children) != 1 {
		t.Errorf("cur.children has length %d and not 1", len(cur.children))
	}
	cur, ok = cur.children['∮']
	if !ok {
		t.Error("cur.children doesn't contain ∮")
	}
	if !cur.term {
		t.Error("cur.terminal is (false) and not (true)")
	}
	if len(cur.children) != 0 {
		t.Errorf("cur.children has length %d and not 0", len(cur.children))
	}

	if !trie.Add("Ӝ\n8") {
		t.Error("trie.Add is (false) and not (true)")
	}
	// trie contains: Ӝ∮∮, <empty string>, Ӝ∮, Ӝ<newline>8
	cur = trie
	if !cur.term {
		t.Error("cur.terminal is (false) and not (true)")
	}
	if len(cur.children) != 1 {
		t.Errorf("cur.children has length %d and not 1", len(cur.children))
	}
	cur, ok = cur.children['Ӝ']
	if !ok {
		t.Error("cur.children doesn't contain Ӝ")
	}
	if cur.term {
		t.Error("cur.terminal is (true) and not (false)")
	}
	if len(cur.children) != 2 {
		t.Errorf("cur.children has length %d and not 1", len(cur.children))
	}
	{
		branch, branchok := cur.children['\n']
		if !branchok {
			t.Error("cur.children doesn't contain '\\n'")
		}
		if len(branch.children) != 1 {
			t.Errorf("branch.children has length %d and not 1", len(branch.children))
		}
		if branch.term {
			t.Error("branch.terminal is (true) and not (false)")
		}
		branch, branchok = branch.children['8']
		if !branchok {
			t.Error("branch.children doesn't contain '8'")
		}
		if len(branch.children) != 0 {
			t.Errorf("branch.children has length %d and not 0", len(branch.children))
		}
		if !branch.term {
			t.Error("branch.terminal is (false) and not (true)")
		}
	}
	cur, ok = cur.children['∮']
	if !ok {
		t.Error("cur.children doesn't contain ∮")
	}
	if !cur.term {
		t.Error("cur.terminal is (false) and not (true)")
	}
	if len(cur.children) != 1 {
		t.Errorf("cur.children has length %d and not 1", len(cur.children))
	}
	cur, ok = cur.children['∮']
	if !ok {
		t.Error("cur.children doesn't contain ∮")
	}
	if !cur.term {
		t.Error("cur.terminal is (false) and not (true)")
	}
	if len(cur.children) != 0 {
		t.Errorf("cur.children has length %d and not 0", len(cur.children))
	}
}
