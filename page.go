package jottit

import (
	"net/url"
	"strconv"
)

type Page struct {
	*url.URL
}

func NewPage(urlstr string) (page *Page, err error) {
	url, err := url.Parse(urlstr)
	if err != nil {
		return
	}

	page = &Page{url}

	return
}

func (page *Page) GetRevisions() (revisions []*Revision, err error) {
	lastRevisionId, err := page.getLastRevisionId()
	if err != nil {
		return
	}

	for i := 0; i < lastRevisionId; i++ {
		revisions = append(revisions, NewRevision(page, i+1))
	}

	return
}

func (page *Page) getLastRevisionId() (id int, err error) {
	relativeUrl, err := url.Parse("?m=history")

	doc, err := url2doc(page.ResolveReference(relativeUrl))
	if err != nil {
		return
	}

	items, err := doc.Search("//input[@name=\"r\"]")
	if err != nil {
		return
	}

	if len(items) == 0 {
		id = 1
	} else {
		id, _ = strconv.Atoi(items[0].Attributes()["value"].String())
	}

	return
}
