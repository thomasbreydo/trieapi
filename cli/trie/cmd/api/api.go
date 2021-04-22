package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const apiURL = "https://trieapi.uk.r.appspot.com/api/v1" // no trailing slash

// WithWord gets https://[api.com]/cmd/word
func WithWord(cmd string, word string) (string, int, error) {
	res, err := http.Get(fmt.Sprintf("%s/%s/%s", apiURL, cmd, word))
	if err != nil {
		return "", 0, err
	}
	txt, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", 0, err
	}
	code := res.StatusCode
	err = res.Body.Close()
	if err != nil {
		return "", 0, err
	}
	return string(txt), code, nil
}

// WithoutWord gets https://[api.com]/cmd/
func WithoutWord(cmd string) (string, int, error) {
	return WithWord(cmd, "")
}
