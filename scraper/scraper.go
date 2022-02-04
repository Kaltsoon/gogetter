package scraper

import (
	"fmt"
	"strings"
	"time"

	"github.com/Kaltsoon/gogetter/link"
	"github.com/Kaltsoon/gogetter/page"
	"github.com/Kaltsoon/gogetter/util"
)

type ScraperResult struct {
	Url         string
	BrokenLinks []*link.Link
}

func CheckLink(link *link.Link, targetPage *page.Page, brokenLinksChannel chan *link.Link) {
	if link.IsBroken(targetPage.Url) {
		brokenLinksChannel <- link
	} else {
		brokenLinksChannel <- nil
	}
}

func ScrapePage(targetPage *page.Page) []*link.Link {
	pageLinks := targetPage.Links()
	brokenLinksChannel := make(chan *link.Link, len(pageLinks))

	for _, link := range pageLinks {
		go CheckLink(link, targetPage, brokenLinksChannel)
	}

	brokenLinks := []*link.Link{}

	for i := 0; i < cap(brokenLinksChannel); i++ {
		value := <-brokenLinksChannel

		if value != nil {
			brokenLinks = append(brokenLinks, value)
		}
	}

	return brokenLinks
}

func ShouldScrapeLink(link *link.Link, targetPage *page.Page, visitedUrls map[string]bool) bool {
	linkUrl, err := link.Url(targetPage.Url)

	if err != nil {
		return false
	}

	if UrlIsVisited(linkUrl, visitedUrls) {
		return false
	}

	return !link.IsBroken(targetPage.Url) && targetPage.IsInternalLink(*link)
}

func NormalizeUrl(url string) string {
	withoutFragment := util.GetUrlWithoutFragment(url)

	normalizedUrl := withoutFragment

	if strings.HasSuffix(withoutFragment, "/") {
		normalizedUrl = normalizedUrl[:len(normalizedUrl)-1]
	}

	return normalizedUrl
}

func UrlIsVisited(url string, visitedUrls map[string]bool) bool {
	_, isVisited := visitedUrls[NormalizeUrl(url)]

	return isVisited
}

func SetUrlVisited(url string, visitedUrls map[string]bool) {
	visitedUrls[NormalizeUrl(url)] = true
}

func RecursiveGetBrokenLinks(targetPage *page.Page, depth int, maxDepth int, visitedUrls map[string]bool) []ScraperResult {
	SetUrlVisited(targetPage.Url, visitedUrls)

	scraperResults := []ScraperResult{}
	pageLinks := targetPage.Links()

	fmt.Printf("\nGoing through %d links at %s", len(pageLinks), targetPage.Url)

	now := time.Now().Unix()

	brokenLinks := ScrapePage(targetPage)

	fmt.Printf("\nDone in %d seconds", time.Now().Unix()-now)

	if len(brokenLinks) > 0 {
		fmt.Printf("\nFound %d broken links", len(brokenLinks))

		scraperResults = append(
			scraperResults,
			ScraperResult{Url: targetPage.Url, BrokenLinks: brokenLinks},
		)
	}

	if depth >= maxDepth {
		return scraperResults
	}

	for _, link := range pageLinks {
		if ShouldScrapeLink(link, targetPage, visitedUrls) {
			linkUrl, _ := link.Url(targetPage.Url)

			normalizedUrl := util.GetUrlWithoutFragment(linkUrl)

			linkScraperResults := RecursiveGetBrokenLinks(
				page.NewFromUrl(normalizedUrl),
				depth+1,
				maxDepth,
				visitedUrls,
			)

			scraperResults = append(
				scraperResults,
				linkScraperResults...,
			)
		}
	}

	return scraperResults
}

func GetBrokenLinks(url string, maxDepth int) []ScraperResult {
	normalizedUrl := util.GetUrlWithoutFragment(url)
	targetPage := page.NewFromUrl(normalizedUrl)
	visitedUrls := make(map[string]bool)

	return RecursiveGetBrokenLinks(targetPage, 0, maxDepth, visitedUrls)
}
