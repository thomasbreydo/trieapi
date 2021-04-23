package main

import (
	"bytes"
	"io/ioutil"
	"strings"
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
	go main() // serve locally for testing purposes
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

	// 0
	cmd.Reset()
	cmd.Root.SetArgs([]string{"add", "--word", ""})
	if err := cmd.Root.Execute(); err != nil {
		t.Error(err)
	}
	check("Keyword () added", buf, t)
	// 1
	cmd.Reset()
	cmd.Root.SetArgs([]string{"add", "--word", "h"})
	if err := cmd.Root.Execute(); err != nil {
		t.Error(err)
	}
	check("Keyword (h) added", buf, t)
	// 2
	cmd.Reset()
	cmd.Root.SetArgs([]string{"add", "--word", ""})
	if err := cmd.Root.Execute(); err != nil {
		t.Error(err)
	}
	check("Keyword () present", buf, t)
	// 3
	cmd.Reset()
	cmd.Root.SetArgs([]string{"delete", "--word", "test"})
	if err := cmd.Root.Execute(); err != nil {
		t.Error(err)
	}
	check("Keyword (test) missing", buf, t)
	// 4
	cmd.Reset()
	cmd.Root.SetArgs([]string{"complete", "-p", ""})
	if err := cmd.Root.Execute(); err != nil {
		t.Error(err)
	}
	check("\nh", buf, t)
	// 5
	cmd.Reset()
	cmd.Root.SetArgs([]string{"search", "--word", "test"})
	if err := cmd.Root.Execute(); err != nil {
		t.Error(err)
	}
	check("Keyword (test) not found", buf, t)
	// 6
	cmd.Reset()
	cmd.Root.SetArgs([]string{"search", "--word", ""})
	if err := cmd.Root.Execute(); err != nil {
		t.Error(err)
	}
	check("Keyword () found", buf, t)
	// 7
	cmd.Reset()
	cmd.Root.SetArgs([]string{"add", "--word", "hello"})
	if err := cmd.Root.Execute(); err != nil {
		t.Error(err)
	}
	check("Keyword (hello) added", buf, t)
	// 8
	cmd.Reset()
	cmd.Root.SetArgs([]string{"add", "--word", "helli"})
	if err := cmd.Root.Execute(); err != nil {
		t.Error(err)
	}
	check("Keyword (helli) added", buf, t)
	// 9
	cmd.Reset()
	cmd.Root.SetArgs([]string{"complete", "--prefix", "hel"})
	if err := cmd.Root.Execute(); err != nil {
		t.Error(err)
	}
	b, err := ioutil.ReadAll(buf)
	if err != nil {
		t.Error(err)
	}
	out := strings.Split(string(b), "\n")
	if out[0] == "hello" && out[1] != "helli" || out[0] == "helli" && out[1] != "hello" {
		t.Errorf("desired: hello\nhelli\nactual:%s", out)
	}
}
