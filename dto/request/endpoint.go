package dto

type Endpoint struct {
	Path            string          `json:"path"`
	Description     string          `json:"description"`
	ExampleRequest  string          `json:"exampleRequest"`
	ExampleResponse string          `json:"exampleResponse"`
	RequestParams   []RequestParam  `json:"requestParams"`
	ResponseParams  []ResponseParam `json:"responseParams"`
}
