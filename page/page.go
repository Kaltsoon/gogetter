package page

import (
	"net/http"
	"strings"

	"github.com/Kaltsoon/gogetter/link"
	"github.com/Kaltsoon/gogetter/util"
	"github.com/PuerkitoBio/goquery"
)

type Page struct {
	Url      string
	Document *goquery.Document
}

func New(url string, document *goquery.Document) *Page {
	return &Page{Url: url, Document: document}
}

func NewFromUrl(url string) *Page {
	res, err := http.Get(url)

	if err != nil || res.StatusCode != 200 {
		return New(url, nil)
	}

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)

	if err != nil {
		return New(url, nil)
	}

	return New(url, doc)
}

func (page Page) Links() []*link.Link {
	links := []*link.Link{}

	if page.Document == nil {
		return links
	}

	linksSelection := page.Document.Find("a")

	if linksSelection == nil {
		return links
	}

	linksSelection.Each(func(index int, selection *goquery.Selection) {
		link := link.New(selection)
		links = append(links, link)
	})

	return links
}

func (page Page) BaseUrl() (string, error) {
	return util.GetBaseUrl(page.Url)
}

func (page Page) IsInternalLink(link link.Link) bool {
	linkUrl, err := link.Url(page.Url)

	if err != nil {
		return false
	}

	pageBaseUrl, err := page.BaseUrl()

	if err != nil {
		return false
	}

	return strings.HasPrefix(linkUrl, pageBaseUrl)
}

func (page Page) String() string {
	return page.Url
}
