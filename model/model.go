package model

// TextResp represents successful response for the testText request.
type TextResp struct {
	Text   string           `json:"text"`
	Topics map[string]int64 `json:"topics"`
}

// IncorrectResp represents generic incorrect response.
type IncorrectResp struct {
	Reason string `json:"reason"`
}

// TextCase represents request for the test-text endpoint.
type TextCase struct {
	Text string `json:"text"`
}
