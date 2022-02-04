package link

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"github.com/Kaltsoon/gogetter/util"
	"github.com/PuerkitoBio/goquery"
)

type Link struct {
	Selection *goquery.Selection
}

type LinkJSON struct {
	Text string
	Href string
}

func New(selection *goquery.Selection) *Link {
	return &Link{Selection: selection}
}

func NewFromHtml(html string) *Link {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))

	var anchor *goquery.Selection

	if err != nil {
		anchor = nil
	} else {
		anchor = doc.Find("a")
	}

	return New(anchor)
}

func (link Link) Text() string {
	if link.Selection == nil {
		return ""
	}

	return link.Selection.Text()
}

func (link Link) Href() string {
	if link.Selection == nil {
		return ""
	}

	return link.Selection.AttrOr("href", "")
}

func (link Link) Url(pageUrl string) (string, error) {
	href := link.Href()

	if strings.HasPrefix(href, "mailto:") {
		return "", fmt.Errorf("can not resolve URL for mailto link %s", href)
	}

	r, _ := regexp.Compile("^https?://")

	if r.MatchString(href) {
		return href, nil
	}

	pageBaseUrl, err := util.GetBaseUrl(pageUrl)

	if err != nil {
		return "", err
	}

	pathPrefix := "/"

	if strings.HasPrefix(href, "/") {
		pathPrefix = ""
	}

	return fmt.Sprintf("%s%s%s", pageBaseUrl, pathPrefix, href), nil
}

func (link Link) IsBroken(pageUrl string) bool {
	if strings.HasPrefix(link.Href(), "mailto:") {
		return false
	}

	linkUrl, err := link.Url(pageUrl)

	if err != nil {
		return true
	}

	isOk, _ := util.UrlIsOk(linkUrl)

	return !isOk
}

func (link Link) String() string {
	return fmt.Sprintf("%s: %s", link.Text(), link.Href())
}

func (link Link) MarshalJSON() ([]byte, error) {
	return json.Marshal(LinkJSON{Text: link.Text(), Href: link.Href()})
}
