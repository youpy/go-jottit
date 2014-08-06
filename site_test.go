package jottit

import (
	"github.com/coocood/assrt"
	"testing"
)

func TestNewSite(t *testing.T) {
	assert := assrt.NewAssert(t)

	site, err := NewSite("http://youpy.jottit.com/")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal("youpy.jottit.com", site.URL.Host)
	assert.Equal("youpy.jottit.com", site.Host)
}

func TestGetPages(t *testing.T) {
	assert := assrt.NewAssert(t)

	site, err := NewSite("http://youpy.jottit.com/")
	if err != nil {
		t.Fatal(err)
	}

	pages, err := site.GetPages()
	if err != nil {
		t.Fatal(err)
	}

	assert.True(len(pages) > 0)

	page := pages[0]

	assert.Equal("youpy.jottit.com", page.Host)
	assert.Equal("/home", page.Path)
}
