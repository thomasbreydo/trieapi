package main

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/thomasbreydo/trieapi/cli/trie/cmd"
	"github.com/thomasbreydo/trieapi/cli/trie/cmd/api"
)

func check(desired string, buf *bytes.Buffer, t *testing.T) {
	b, err := ioutil.ReadAll(buf)
	if err != nil {
		t.Error(err)
	}
	if actual := string(b); desired != actual {
		t.Errorf("desired: %s\nactual: %s", desired, actual)
	}
}

func TestServer(t *testing.T) {
	go main()
	api.URL = "http://localhost:8080/api/v1"
	buf := bytes.NewBuffer(nil)
	cmd.Reset()
	cmd.Root.SetOut(buf)

	cmd.Reset()
	cmd.Root.SetArgs([]string{"display", "--json"})
	if err := cmd.Root.Execute(); err != nil {
		t.Error(err)
	}
	check("[]", buf, t)

	cmd.Reset()
	cmd.Root.SetArgs([]string{"add", "-w", "test"})
	if err := cmd.Root.Execute(); err != nil {
		t.Error(err)
	}
	check("Keyword (test) added", buf, t)

	cmd.Reset()
	cmd.Root.SetArgs([]string{"display", "--json"})
	if err := cmd.Root.Execute(); err != nil {
		t.Error(err)
	}
	check("[\"test\"]", buf, t)

	cmd.Reset()
	cmd.Root.SetArgs([]string{"display"})
	if err := cmd.Root.Execute(); err != nil {
		t.Error(err)
	}
	check("test", buf, t)

	cmd.Reset()
	cmd.Root.SetArgs([]string{"search", "-w", "test"})
	if err := cmd.Root.Execute(); err != nil {
		t.Error(err)
	}
	check("Keyword (test) found", buf, t)

	cmd.Reset()
	cmd.Root.SetArgs([]string{"add", "--word", "test"})
	if err := cmd.Root.Execute(); err != nil {
		t.Error(err)
	}
	check("Keyword (test) present", buf, t)

	cmd.Reset()
	cmd.Root.SetArgs([]string{"delete", "--word", "test"})
	if err := cmd.Root.Execute(); err != nil {
		t.Error(err)
	}
	check("Keyword (test) deleted", buf, t)

	cmd.Reset()
	cmd.Root.SetArgs([]string{"delete", "--word", "test"})
	if err := cmd.Root.Execute(); err != nil {
		t.Error(err)
	}
	check("Keyword (test) missing", buf, t)
}
