package app

// DFRequest struct
type DFRequest struct {
	Result struct {
		Action        string            `json:"action"`
		Parameters    map[string]string `json:"parameters"`
		ResolvedQuery string            `json:"resolvedQuery"`
	} `json:"result"`
	OriginalRequest DFOriginalRequest `json:"originalRequest"`
}

// DFOriginalRequest struct
type DFOriginalRequest struct {
}

// DFResponse struct
type DFResponse struct {
	Speech string         `json:"speech"`
	Data   DFResponseData `json:"data"`
}

// DFResponseData struct
type DFResponseData struct {
	Google DFResponseGoogle `json:"google"`
}

// DFResponseGoogle struct
type DFResponseGoogle struct {
	ExpectResponse bool `json:"expect_user_response"`
	IsSsml         bool `json:"isSsml"`
}
