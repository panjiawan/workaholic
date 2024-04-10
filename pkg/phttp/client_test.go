package phttp

import "testing"

func TestPost(t *testing.T) {
	var (
		uri   = ""
		param = map[string]string{}
	)
	body, err := Post(uri, param, nil)
	if err != nil {
		t.Error(err)
	}

	t.Log(string(body))
}

func TestGet(t *testing.T) {
	var (
		uri    = ""
		param  = map[string]string{}
		header = map[string]string{}
	)
	body, err := Get(uri, param, header)
	if err != nil {
		t.Error(err)
	}

	t.Log(string(body))
}
