package dto

type Endpoint struct {
	Description    string          `json:"description"`
	RequestParams  []RequestParam  `json:"requestParams"`
	ResponseParams []ResponseParam `json:"responseParams"`
}
