package jottitit

import (
	"net/url"
)

type Site struct {
	*url.URL
}

func NewSite(urlstr string) (site *Site, err error) {
	url, err := url.Parse(urlstr)
	if err != nil {
		return
	}

	site = &Site{url}

	return
}

func (site *Site) GetPages() (pages []*Page, err error) {
	firstPageRelativeURL, err := url.Parse("/home")
	if err != nil {
		return
	}

	page, err := NewPage(site.ResolveReference(firstPageRelativeURL).String())
	if err != nil {
		return
	}

	pages = append(pages, page)

	doc, err := url2doc(site.URL)
	if err != nil {
		return
	}

	items, err := doc.Search("//ul[@id=\"pages\"]/li/a")

	if err != nil {
		return
	}

	for _, item := range items {
		page, err = NewPage(item.Attributes()["href"].String())
		if err != nil {
			return nil, err
		}

		pages = append(pages, page)
	}

	return
}
