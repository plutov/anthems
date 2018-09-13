package app

import (
	"fmt"
	"net/http"
	"testing"
	"time"
)

func TestCountriesFileExist(t *testing.T) {
	for _, c := range supportedCountries {
		country := c

		t.Run(country, func(subtest *testing.T) {
			subtest.Parallel()

			fileName, replace := replaceMap[country]
			if !replace {
				fileName = getFileName(country)
			}

			url := fmt.Sprintf("%s/%s.mp3", sourceDir, fileName)

			httpClient := http.Client{
				Timeout: time.Duration(10 * time.Second),
			}

			resp, err := httpClient.Head(url)
			if err != nil || resp == nil || resp.StatusCode >= http.StatusBadRequest {
				subtest.Fatalf("url doesn't exist: %s, err: %v", url, err)
			}
			resp.Body.Close()
		})
	}
}
