package main

import (
	"flag"

	"github.com/Kaltsoon/gogetter/reporter"
	"github.com/Kaltsoon/gogetter/scraper"
)

func main() {
	url := flag.String("url", "", "The URL of a web page to start the search of the broken links")
	maxDepth := flag.Int("maxdepth", 5, "The maximum depth of the recursive search")
	reportFile := flag.String("out", "report.json", "Name of the output file where the information about the broken links is written in JSON")

	flag.Parse()

	if *url == "" {
		panic("the url flag is required")
	}

	results := scraper.GetBrokenLinks(*url, *maxDepth)

	reporter.WriteReport(results, *reportFile)
}
