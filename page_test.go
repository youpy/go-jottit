package jottit

import (
	"github.com/coocood/assrt"
	"testing"
)

func TestNewPage(t *testing.T) {
	assert := assrt.NewAssert(t)

	page, err := NewPage("http://youpy.jottit.com/1")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal("youpy.jottit.com", page.URL.Host)
	assert.Equal("youpy.jottit.com", page.Host)
}

func TestGetRevisions(t *testing.T) {
	assert := assrt.NewAssert(t)

	page, err := NewPage("http://youpy.jottit.com/1")
	if err != nil {
		t.Fatal(err)
	}

	revisions, err := page.GetRevisions()
	if err != nil {
		t.Fatal(err)
	}

	assert.True(len(revisions) > 0)
	assert.True(revisions[0].Id > 0)
	assert.Equal(13, revisions[len(revisions)-1].Id)
}
