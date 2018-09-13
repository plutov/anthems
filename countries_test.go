package app

import (
	"fmt"
	"net/http"
	"net/url"
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

			url := fmt.Sprintf("%s/%s.mp3", sourceDir, url.QueryEscape(fileName))

			httpClient := http.Client{
				Timeout: time.Duration(5 * time.Second),
			}

			resp, err := httpClient.Get(url)
			if err != nil || resp == nil || resp.StatusCode >= http.StatusBadRequest {
				subtest.Fatalf("unable to load url %s, err: %v", url, err)
			}
			resp.Body.Close()
		})
	}
}
