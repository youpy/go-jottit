package jottit

import (
	"github.com/coocood/assrt"
	"testing"
)

func TestNewRevision(t *testing.T) {
	assert := assrt.NewAssert(t)

	page, err := NewPage("http://youpy.jottit.com/1")
	if err != nil {
		t.Fatal(err)
	}

	revision := NewRevision(page, 1)

	assert.Equal(1, revision.Id)
}

func testGetFullURL(t *testing.T) {
	assert := assrt.NewAssert(t)

	page, err := NewPage("http://youpy.jottit.com/1")
	if err != nil {
		t.Fatal(err)
	}

	revision := NewRevision(page, 1)

	fullURL, err := revision.GetFullURL()
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal("http://youpy.jottit.com/1?r=1", fullURL)
}

func TestGetConent(t *testing.T) {
	assert := assrt.NewAssert(t)

	page, err := NewPage("http://youpy.jottit.com/1")
	if err != nil {
		t.Fatal(err)
	}

	revision := NewRevision(page, 1)

	content, err := revision.GetContent()
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal("* あ\n* ん\n", content)
}

func TestGetPostTime(t *testing.T) {
	assert := assrt.NewAssert(t)

	page, err := NewPage("http://youpy.jottit.com/1")
	if err != nil {
		t.Fatal(err)
	}

	revision := NewRevision(page, 1)

	postTime, err := revision.GetPostTime()
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal("2010-05-12 00:00:00 +0000 UTC", postTime.String())
}
