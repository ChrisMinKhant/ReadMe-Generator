package dto

type RequestParam struct {
	Field       string       `json:"field"`
	Type        string       `json:"type"`
	Validations []Validation `json:"validations"`
	Description string       `json:"description"`
}
