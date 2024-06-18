package dto

type RequestParam struct {
	Field       string       `json:"field"`
	Validations []Validation `json:"validations"`
	Description string       `json:"description"`
}
