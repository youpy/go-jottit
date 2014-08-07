package jottit

import (
	"fmt"
	"net/url"
	"regexp"
	"strconv"
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
	relativeUrl, err := revision.Parse(fmt.Sprintf("?m=diff&r=%d&r=%d", revision.Id, revision.Id))
	if err != nil {
		return
	}

	diffUrl := revision.Page.ResolveReference(relativeUrl)

	doc, err := url2doc(diffUrl)
	if err != nil {
		return
	}

	items, err := doc.Search("//div[@id=\"diff\"]/code")
	if err != nil {
		return
	}

	if len(items) == 0 {
		err = fmt.Errorf("can't find content for %s\n", revision.Page)
		return
	}

	content = items[0].InnerHtml()

	re, err := regexp.Compile("<br[^/]*>")
	if err != nil {
		return
	}

	content = re.ReplaceAllString(content, "")
	if len(content) > 0 {
		content = content[1:]
	}

	return
}
