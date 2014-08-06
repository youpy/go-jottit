package jottitit

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

	content, _ := revision.GetContent()

	assert.True(len(content) > 0)
}
