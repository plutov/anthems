package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

func init() {
	http.HandleFunc("/", handle)
}

const (
	s3Bucket   = "https://s3.amazonaws.com/actions-on-google-anthems"
	welcomeMsg = "Hi! I can play a national anthem of any country, which country do you want?"
	noMsg      = "Sure, thanks for talking to Anthems. Goodbye!"
	errMsg     = "Sorry, I didn't get that. Can you please try again?"
)

var (
	supportedCountries = map[string]string{
		"belarus":                  "ogg",
		"ukraine":                  "oga",
		"united states of america": "ogg",
	}
)

func handle(w http.ResponseWriter, r *http.Request) {
	dfReq := DFRequest{}
	dfErr := json.NewDecoder(r.Body).Decode(&dfReq)

	if dfErr == nil && dfReq.Result.Action == "input.welcome" {
		json.NewEncoder(w).Encode(DFResponse{
			Speech: "<speak>" + welcomeMsg + "</speak>",
			Data: DFResponseData{
				Google: DFResponseGoogle{
					IsSsml: true,
				},
			},
		})
		return
	}

	if dfErr == nil && dfReq.Result.Action == "input.country" {
		country, ok := dfReq.Result.Parameters["geo-country"]
		if !ok || len(country) == 0 {
			country = dfReq.Result.ResolvedQuery
		}

		country = strings.ToLower(country)
		ext, supported := supportedCountries[country]

		if len(country) == 0 || !supported {
			returnAPIErrorMessage(w)
			return
		}

		json.NewEncoder(w).Encode(DFResponse{
			Speech: fmt.Sprintf(`<speak>Cool, %s. <audio src="%s/%s.%s"></audio>. That's it, do you want to listen another country anthem?</speak>`, country, s3Bucket, url.QueryEscape(country), ext),
			Data: DFResponseData{
				Google: DFResponseGoogle{
					IsSsml: true,
				},
			},
		})
		return
	}

	if dfErr == nil && dfReq.Result.Action == "input.no" {
		json.NewEncoder(w).Encode(DFResponse{
			Speech: "<speak>" + noMsg + "</speak>",
			Data: DFResponseData{
				Google: DFResponseGoogle{
					IsSsml: true,
				},
			},
		})
		return
	}

	returnAPIErrorMessage(w)
}

func returnAPIErrorMessage(w http.ResponseWriter) {
	json.NewEncoder(w).Encode(DFResponse{
		Speech: "<speak>" + errMsg + "</speak>",
		Data: DFResponseData{
			Google: DFResponseGoogle{
				IsSsml: true,
			},
		},
	})
}

func handleGetAction(w http.ResponseWriter, r *http.Request, dfReq DFRequest) {

}
