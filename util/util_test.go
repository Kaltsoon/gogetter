package util

import "testing"

func TestGetBaseUrl(t *testing.T) {
	baseUrl, _ := GetBaseUrl("https://foo.bar/john?doe=lorem#ipsum")

	if baseUrl != "https://foo.bar" {
		t.Errorf("expected base url to be https://foo.bar but was %s", baseUrl)
	}
}

func TestGetUrlWithoutFragment(t *testing.T) {
	url := GetUrlWithoutFragment("https://foo.bar/john?doe=lorem#ipsum")

	if url != "https://foo.bar/john?doe=lorem" {
		t.Errorf("expected base url to be https://foo.bar/john?doe=lorem but was %s", url)
	}
}
