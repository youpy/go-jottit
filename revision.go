package jottit

import (
	"fmt"
	"net/url"
	"regexp"
	"strconv"
	"time"
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

	content = items[0].Content()
	if len(content) > 0 {
		content = content[1:]
	}

	return
}

func (revision *Revision) GetPostTime() (postTime time.Time, err error) {
	url, err := revision.GetFullURL()
	if err != nil {
		return
	}

	doc, err := url2doc(url)
	if err != nil {
		return
	}

	items, err := doc.Search("//div[@id=\"left\"]//strong")
	if err != nil {
		return
	}

	if len(items) == 0 {
		err = fmt.Errorf("can't find content for %s\n", revision.Page)
		return
	}

	timeString := items[0].InnerHtml()

	matched, err := regexp.MatchString(" \\d{4}$", timeString)
	if err != nil {
		return
	}
	if !matched {
		timeString += fmt.Sprintf(", %d", time.Now().Year())
	}

	re := regexp.MustCompile("(\\w{3})\\w* *(\\d+), \\d{2}(\\d{2})")
	timeString = re.ReplaceAllString(timeString, "${2} ${1} ${3} 00:00 UTC")

	postTime, err = time.Parse("2 Jan 06 15:04 MST", timeString)
	if err != nil {
		return
	}

	return
}
