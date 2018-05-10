package security

import (
	"errors"
	"net/http"
	"regexp"
)

var validPath = regexp.MustCompile("^/(edit|view|save)/([a-zA-Z0-9]+)$")

func CheckPath(w http.ResponseWriter, r *http.Request) (string, error) {

	m := validPath.FindStringSubmatch(r.URL.Path)
	if nil == m {
		http.NotFound(w, r)
		return "", errors.New("Invalid Page Title")
	}
	return m[2], nil
}
