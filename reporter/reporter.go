package reporter

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Kaltsoon/gogetter/scraper"
)

func WriteReport(scrapperResults []scraper.ScraperResult, fileName string) (bool, error) {
	data, err := json.MarshalIndent(scrapperResults, "", "  ")

	if err != nil {
		return false, err
	}

	filePath := fmt.Sprintf("data/%s", fileName)

	file, err := os.Create(filePath)

	if err != nil {
		return false, err
	}

	defer file.Close()

	_, err = file.Write(data)

	if err != nil {
		return false, err
	}

	return true, nil
}
