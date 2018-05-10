package page

import (
	"testing"
)

func TestSave(t *testing.T) {

	// save
	page := &Page{Title: "yao", Body: []byte("hello world!")}
	page.Save()

	// read-yao
	readPage, err := LoadPage("yao")
	if nil != err {
		t.Error(err)
	} else {
		t.Log(string(readPage.Body))
	}

	// read-404
	readPage, err = LoadPage("404")
	if nil != err {
		t.Error(err)
	} else {
		t.Log(string(readPage.Body))
	}
}
