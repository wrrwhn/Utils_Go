package page

import (
	"fmt"
	"io/ioutil"
	"log"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) Save() error {
	filename := getLocalPath(p.Title)
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func LoadPage(title string) (*Page, error) {
	filename := getLocalPath(title)
	body, err := ioutil.ReadFile(filename)
	if nil != err {
		log.Printf("fail to load page[%v]", filename)
		return nil, err
	}

	return &Page{Title: title, Body: body}, nil
}

func getLocalPath(title string) string {
	return fmt.Sprintf("data/%s.txt", title)
}
