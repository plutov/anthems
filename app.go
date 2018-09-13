package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"unicode"
)

func init() {
	http.HandleFunc("/", handle)
}

const (
	sourceDir  = "https://anthemworld.com/~anthemwo/themesongs"
	welcomeMsg = "Hi! I can play a national anthem of any country, which country do you want?"
	noMsg      = "Sure, thanks for talking to Anthems. Goodbye!"
	yesMsg     = "Great, which country do you want?"
	errMsg     = "Sorry, I didn't get that. Can you please try again?"
)

func handle(w http.ResponseWriter, r *http.Request) {
	dfReq := DFRequest{}
	dfErr := json.NewDecoder(r.Body).Decode(&dfReq)

	if dfErr == nil && dfReq.Result.Action == "input.welcome" {
		returnSSMLResponse(w, welcomeMsg, true)
		return
	}

	if dfErr == nil && dfReq.Result.Action == "input.country" {
		country, ok := dfReq.Result.Parameters["geo-country"]
		if !ok || len(country) == 0 {
			country = dfReq.Result.ResolvedQuery
		}

		country = strings.ToLower(country)

		if len(country) == 0 || !stringInSlice(country, supportedCountries) {
			returnAPIErrorMessage(w)
			return
		}

		fileName, replace := replaceMap[country]
		if !replace {
			fileName = getFileName(country)
		}

		returnSSMLResponse(w, fmt.Sprintf(`Cool, %s. <audio src="%s/%s.mp3">Sorry, couldn't find the %s anthem.</audio> Do you want to listen another country anthem?`, country, sourceDir, fileName, country), true)
		return
	}

	if dfErr == nil && dfReq.Result.Action == "country.country-no" {
		returnSSMLResponse(w, noMsg, false)
		return
	}

	if dfErr == nil && dfReq.Result.Action == "country.country-yes" {
		returnSSMLResponse(w, yesMsg, true)
		return
	}

	returnAPIErrorMessage(w)
}

func returnAPIErrorMessage(w http.ResponseWriter) {
	returnSSMLResponse(w, errMsg, true)
}

func returnSSMLResponse(w http.ResponseWriter, msg string, expectResp bool) {
	json.NewEncoder(w).Encode(DFResponse{
		Speech: "<speak>" + msg + "</speak>",
		Data: DFResponseData{
			Google: DFResponseGoogle{
				ExpectResponse: expectResp,
				IsSsml:         true,
			},
		},
	})
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// 1. Capitalize 1st letter in each word
func getFileName(country string) string {
	var words = strings.Split(country, " ")
	var capWords []string
	for _, w := range words {
		var capWord string
		if w == "and" {
			capWord = "and"
		} else {
			for _, v := range w {
				u := string(unicode.ToUpper(v))
				capWord = u + w[len(u):]
				break
			}
		}

		capWords = append(capWords, capWord)
	}

	return strings.Join(capWords, " ")
}
