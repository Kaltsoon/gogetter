package util

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"time"

	lru "github.com/hashicorp/golang-lru"
)

var client = &http.Client{
	Timeout: time.Second * 5,
}

var htmlCache, _ = lru.New(200)
var pingCache, _ = lru.New(200)

func UrlIsOk(url string) (bool, error) {
	ping, inCache := pingCache.Get(url)

	if inCache {
		return ping.(bool), nil
	}

	res, err := client.Get(url)

	if err != nil {
		pingCache.Add(url, false)
		return false, err
	}

	defer res.Body.Close()

	if res.StatusCode == 401 || res.StatusCode >= 500 {
		pingCache.Add(url, false)
		return false, fmt.Errorf("received bad status code %d from URL %s", res.StatusCode, url)
	}

	pingCache.Add(url, true)

	return true, nil
}

func GetUrlHtml(url string) (string, error) {
	html, inCache := htmlCache.Get(url)

	if inCache {
		return html.(string), nil
	}

	res, err := client.Get(url)

	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		return "", err
	}

	bodyString := string(body)

	htmlCache.Add(url, bodyString)

	return bodyString, nil
}

func GetBaseUrl(rawUrl string) (string, error) {
	parsedUrl, err := url.Parse(rawUrl)

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s://%s", parsedUrl.Scheme, parsedUrl.Host), nil
}

func GetUrlWithoutFragment(rawUrl string) string {
	r, _ := regexp.Compile("(.+)#")

	matches := r.FindStringSubmatch(rawUrl)

	if len(matches) > 1 {
		return matches[1]
	}

	return rawUrl
}
