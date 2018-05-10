package security

import (
	"net/http"
	"net/url"
	"testing"
)

func TestCheckPath(t *testing.T) {

	w := http.ResponseWriter{}
	r := &http.Request{URL: &url.URL{Path: "/view/yao"}}
	title, err := CheckPath(w, r)
	if nil != err {
		t.Log(title)
	} else {
		t.Error()
	}

}
