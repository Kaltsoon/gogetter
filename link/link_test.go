package link

import "testing"

func TestNewFromHtmlWithTextAndHref(t *testing.T) {
	link := NewFromHtml("<a href=\"https://foo.bar\">foo</a>")

	if link.Href() != "https://foo.bar" {
		t.Errorf("expected href to be https://foo.bar but was %s", link.Href())
	}

	if link.Text() != "foo" {
		t.Errorf("expect text to be foo but was %s", link.Text())
	}
}

func TestNewFromHtmlWithoutHref(t *testing.T) {
	link := NewFromHtml("<a>foo</a>")

	if link.Href() != "" {
		t.Errorf("expected href to be an empty string but was %s", link.Href())
	}
}

func TestUrlWithUrl(t *testing.T) {
	link := NewFromHtml("<a href=\"https://foo.bar\">foo</a>")

	url, _ := link.Url("https://john.doe")

	if url != "https://foo.bar" {
		t.Errorf("expected url to be https://foo.bar but was %s", url)
	}
}

func TestUrlWithPath(t *testing.T) {
	link := NewFromHtml("<a href=\"/path\">foo</a>")

	url, _ := link.Url("https://foo.bar")

	if url != "https://foo.bar/path" {
		t.Errorf("expected url to be https://foo.bar/path but was %s", url)
	}
}

func TestUrlWithPathMissingSlash(t *testing.T) {
	link := NewFromHtml("<a href=\"path\">foo</a>")

	url, _ := link.Url("https://foo.bar")

	if url != "https://foo.bar/path" {
		t.Errorf("expected url to be https://foo.bar/path but was %s", url)
	}
}
