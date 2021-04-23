package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const URL = "https://trieapi.uk.r.appspot.com/api/v1" // no trailing slash

// withWord gets https://[api.com]/cmd/word
func withWord(cmd, w string) ([]byte, int, error) {
	res, err := http.Get(fmt.Sprintf("%s/%s/%s", URL, cmd, w))
	if err != nil {
		return nil, 0, err
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, 0, err
	}
	code := res.StatusCode
	err = res.Body.Close()
	if err != nil {
		return nil, 0, err
	}
	return body, code, nil
}

// noWord gets https://[api.com]/cmd/
func noWord(cmd string) ([]byte, int, error) {
	return withWord(cmd, "")
}

type foundInfo struct {
	found bool
}

type modInfo struct {
	modified bool
}

func Add(w string) (bool, int, error) {
	dat, code, err := withWord("add", w)
	if err != nil {
		return false, code, err
	}
	var info modInfo
	err = json.Unmarshal(dat, &info)
	if err != nil {
		return false, code, err
	}
	return info.modified, code, nil
}

func Clear() (bool, int, error) {
	dat, code, err := noWord("clear")
	if err != nil {
		return false, code, err
	}
	var info modInfo
	err = json.Unmarshal(dat, &info)
	if err != nil {
		return false, code, err
	}
	return info.modified, code, nil
}

func CompleteJSON(w string) ([]byte, int, error) {
	return withWord("complete", w)
}

func Complete(w string) (string, int, error) {
	dat, code, err := withWord("complete", w)
	if err != nil {
		return "", code, err
	}
	var comps []string
	err = json.Unmarshal(dat, &comps)
	if err != nil {
		return "", code, err
	}
	return strings.Join(comps, "\n"), code, nil
}

func Delete(w string) (bool, int, error) {
	dat, code, err := withWord("delete", w)
	if err != nil {
		return false, code, err
	}
	var info modInfo
	err = json.Unmarshal(dat, &info)
	if err != nil {
		return false, code, err
	}
	return info.modified, code, nil
}

func DisplayJSON() ([]byte, int, error) {
	return noWord("display")
}

func Display() (string, int, error) {
	dat, code, err := noWord("display")
	if err != nil {
		return "", code, err
	}
	var ww []string
	err = json.Unmarshal(dat, &ww)
	if err != nil {
		return "", code, err
	}
	return strings.Join(ww, "\n"), code, nil
}

func Search(w string) (bool, int, error) {
	dat, code, err := withWord("search", "w")
	if err != nil {
		return false, code, err
	}
	var info foundInfo
	err = json.Unmarshal(dat, &info)
	if err != nil {
		return false, code, err
	}
	return info.found, code, nil
}
