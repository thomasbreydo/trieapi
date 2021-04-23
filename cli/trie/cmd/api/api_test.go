package api

import (
	"encoding/json"
	"testing"
)

func TestInfo(t *testing.T) {
	m, _ := json.Marshal(map[string]bool{"found": true})
	var f foundInfo
	err := json.Unmarshal(m, &f)
	if err != nil {
		t.Error(err)
	}
	if !f.Found {
		t.Error("expected true")
	}
	m, err = json.Marshal(map[string]bool{"found": false})
	if err != nil {
		t.Error(err)
	}
	err = json.Unmarshal(m, &f)
	if err != nil {
		t.Error(err)
	}
	if f.Found {
		t.Error("expected false")
	}
	m, err = json.Marshal(map[string]bool{"modified": true})
	if err != nil {
		t.Error(err)
	}
	var mod modInfo
	err = json.Unmarshal(m, &mod)
	if err != nil {
		t.Error(err)
	}
	if !mod.Mod {
		t.Error("expected true")
	}
	m, err = json.Marshal(map[string]bool{"modified": false})
	if err != nil {
		t.Error(err)
	}
	err = json.Unmarshal(m, &mod)
	if err != nil {
		t.Error(err)
	}
	if mod.Mod {
		t.Error("expected false")
	}
}
