package jottitit

import (
	"fmt"
	"net/url"
	"regexp"
	"strconv"
	"strings"
)

type Revision struct {
	*Page
	Id int
}

func NewRevision(page *Page, id int) (revision *Revision) {
	revision = &Revision{page, id}

	return
}

func (revision *Revision) GetFullURL() (url *url.URL, err error) {
	relativeUrl, err := revision.Parse("?r=" + strconv.Itoa(revision.Id))

	if err != nil {
		return
	}

	url = revision.Page.ResolveReference(relativeUrl)

	return
}

func (revision *Revision) GetContent() (content string, err error) {
	fullURL, err := revision.GetFullURL()
	if err != nil {
		return
	}

	doc, err := url2doc(fullURL)
	if err != nil {
		return
	}

	items, err := doc.Search("//div[@id=\"content\"]")
	if err != nil {
		return
	}

	if len(items) == 0 {
		err = fmt.Errorf("can't find content for %s\n", revision.Page)
		return
	}

	content = items[0].InnerHtml()

	re, _ := regexp.Compile("(?s)^\\s+<form.+?</form>")
	content = re.ReplaceAllString(content, "")

	re, _ = regexp.Compile("(?s)<div id=\"dateline\".+?</div>\\s+$")
	content = re.ReplaceAllString(content, "")

	content = strings.Trim(content, "\n\r\t ")

	return
}
