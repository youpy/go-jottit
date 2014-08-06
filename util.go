package jottitit

import (
	// https://github.com/SciRuby/sciruby/issues/25
	"github.com/moovweb/gokogiri"
	"github.com/moovweb/gokogiri/html"

	"io/ioutil"
	"net/http"
	"net/url"
)

func url2doc(url *url.URL) (doc *html.HtmlDocument, err error) {
	resp, err := http.Get(url.String())
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	doc, err = gokogiri.ParseHtml(body)
	if err != nil {
		return
	}

	return
}
