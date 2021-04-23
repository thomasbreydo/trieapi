package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
	"testing"

	"github.com/thomasbreydo/trieapi/cli/trie/cmd"
	"github.com/thomasbreydo/trieapi/cli/trie/cmd/api"
)

func checkStr(desired string, buf *bytes.Buffer, t *testing.T) {
	b, err := ioutil.ReadAll(buf)
	if err != nil {
		t.Error(err)
	}
	if actual := string(b); desired != actual {
		t.Errorf("desired: %q\nactual: %q", desired, actual)
	}
}

func checkJSON(desired []string, buf *bytes.Buffer, t *testing.T) {
	b, err := ioutil.ReadAll(buf)
	if err != nil {
		t.Error(err)
	}
	var actual []string
	err = json.Unmarshal(b, &actual)
	if err != nil {
		t.Error(err)
	}
	if (actual == nil) != (desired == nil) {
		t.Errorf("desired: %q\nactual: %q", desired, actual)
		return
	}
	if len(actual) != len(desired) {
		t.Errorf("desired: %s\nactual: %s", desired, actual)
		return
	}
	sort.Strings(actual)
	sort.Strings(desired)
	for i, r := range actual {
		if r != desired[i] {
			t.Errorf("desired: %s\nactual: %s", desired, actual)
			return
		}
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
	checkStr("[]", buf, t)

	cmd.Reset()
	cmd.Root.SetArgs([]string{"add", "-w", "test"})
	if err := cmd.Root.Execute(); err != nil {
		t.Error(err)
	}
	checkStr("Keyword (test) added", buf, t)

	cmd.Reset()
	cmd.Root.SetArgs([]string{"display", "--json"})
	if err := cmd.Root.Execute(); err != nil {
		t.Error(err)
	}
	checkStr("[\"test\"]", buf, t)

	cmd.Reset()
	cmd.Root.SetArgs([]string{"display"})
	if err := cmd.Root.Execute(); err != nil {
		t.Error(err)
	}
	checkStr("test", buf, t)

	cmd.Reset()
	cmd.Root.SetArgs([]string{"search", "-w", "test"})
	if err := cmd.Root.Execute(); err != nil {
		t.Error(err)
	}
	checkStr("Keyword (test) found", buf, t)

	cmd.Reset()
	cmd.Root.SetArgs([]string{"add", "--word", "test"})
	if err := cmd.Root.Execute(); err != nil {
		t.Error(err)
	}
	checkStr("Keyword (test) present", buf, t)

	cmd.Reset()
	cmd.Root.SetArgs([]string{"delete", "--word", "test"})
	if err := cmd.Root.Execute(); err != nil {
		t.Error(err)
	}
	checkStr("Keyword (test) deleted", buf, t)

	cmd.Reset()
	cmd.Root.SetArgs([]string{"search", "-w", "test"})
	if err := cmd.Root.Execute(); err != nil {
		t.Error(err)
	}
	checkStr("Keyword (test) not found", buf, t)

	cmd.Reset()
	cmd.Root.SetArgs([]string{"delete", "--word", "test"})
	if err := cmd.Root.Execute(); err != nil {
		t.Error(err)
	}
	checkStr("Keyword (test) missing", buf, t)

	cmd.Reset()
	cmd.Root.SetArgs([]string{"search", "-w", "test"})
	if err := cmd.Root.Execute(); err != nil {
		t.Error(err)
	}
	checkStr("Keyword (test) not found", buf, t)

	cmd.Reset()
	cmd.Root.SetArgs([]string{"search", "-w", ""})
	if err := cmd.Root.Execute(); err != nil {
		t.Error(err)
	}
	checkStr("Keyword () not found", buf, t)

	cmd.Reset()
	cmd.Root.SetArgs([]string{"add", "--word", ""})
	if err := cmd.Root.Execute(); err != nil {
		t.Error(err)
	}
	checkStr("Keyword () added", buf, t)

	cmd.Reset()
	cmd.Root.SetArgs([]string{"search", "-w", ""})
	if err := cmd.Root.Execute(); err != nil {
		t.Error(err)
	}
	checkStr("Keyword () found", buf, t)

	cmd.Reset()
	cmd.Root.SetArgs([]string{"add", "--word", "h"})
	if err := cmd.Root.Execute(); err != nil {
		t.Error(err)
	}
	checkStr("Keyword (h) added", buf, t)

	cmd.Reset()
	cmd.Root.SetArgs([]string{"add", "--word", ""})
	if err := cmd.Root.Execute(); err != nil {
		t.Error(err)
	}
	checkStr("Keyword () present", buf, t)

	cmd.Reset()
	cmd.Root.SetArgs([]string{"delete", "--word", "test"})
	if err := cmd.Root.Execute(); err != nil {
		t.Error(err)
	}
	checkStr("Keyword (test) missing", buf, t)

	cmd.Reset()
	cmd.Root.SetArgs([]string{"complete", "-p", ""})
	if err := cmd.Root.Execute(); err != nil {
		t.Error(err)
	}
	checkStr("\nh", buf, t)

	cmd.Reset()
	cmd.Root.SetArgs([]string{"search", "--word", "test"})
	if err := cmd.Root.Execute(); err != nil {
		t.Error(err)
	}
	checkStr("Keyword (test) not found", buf, t)

	cmd.Reset()
	cmd.Root.SetArgs([]string{"search", "--word", ""})
	if err := cmd.Root.Execute(); err != nil {
		t.Error(err)
	}
	checkStr("Keyword () found", buf, t)

	cmd.Reset()
	cmd.Root.SetArgs([]string{"add", "--word", "hello"})
	if err := cmd.Root.Execute(); err != nil {
		t.Error(err)
	}
	checkStr("Keyword (hello) added", buf, t)

	cmd.Reset()
	cmd.Root.SetArgs([]string{"add", "--word", "helli"})
	if err := cmd.Root.Execute(); err != nil {
		t.Error(err)
	}
	checkStr("Keyword (helli) added", buf, t)

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
	if out[0] == "hello" && out[1] != "helli" ||
		out[0] == "helli" && out[1] != "hello" {
		t.Errorf(
			"desired: hello\nhelli\nor:helli\nhello\nactual:%s", out)
	}

	cmd.Reset()
	cmd.Root.SetArgs([]string{"add", "--word", "hello"})
	if err = cmd.Root.Execute(); err != nil {
		t.Error(err)
	}
	checkStr("Keyword (hello) present", buf, t)

	state := []string{"", "h", "hello", "helli"}
	for i, s := range []string{"*", "fj", "amazon", "amazing", "ample", "apple"} {
		cmd.Reset()
		if i%2 == 0 {
			cmd.Root.SetArgs([]string{"search", "--word", s})
		} else {
			cmd.Root.SetArgs([]string{"search", "-w", s})
		}
		if err = cmd.Root.Execute(); err != nil {
			t.Error(err)
		}
		checkStr(fmt.Sprintf("Keyword (%s) not found", s), buf, t)
		cmd.Reset()
		cmd.Root.SetArgs([]string{"display", "--json"})
		if err = cmd.Root.Execute(); err != nil {
			t.Error(err)
		}
		checkJSON(state, buf, t)
		cmd.Reset()
		if i%2 == 0 {
			cmd.Root.SetArgs([]string{"add", "--word", s})
		} else {
			cmd.Root.SetArgs([]string{"add", "-w", s})
		}
		if err = cmd.Root.Execute(); err != nil {
			t.Error(err)
		}
		checkStr(fmt.Sprintf("Keyword (%s) added", s), buf, t)
		state = append(state, s)
		cmd.Reset()
		cmd.Root.SetArgs([]string{"display", "--json"})
		if err = cmd.Root.Execute(); err != nil {
			t.Error(err)
		}
		checkJSON(state, buf, t)
		cmd.Reset()
		if i%2 == 0 {
			cmd.Root.SetArgs([]string{"search", "--word", s})
		} else {
			cmd.Root.SetArgs([]string{"search", "-w", s})
		}
		if err = cmd.Root.Execute(); err != nil {
			t.Error(err)
		}
		checkStr(fmt.Sprintf("Keyword (%s) found", s), buf, t)
	}
}
