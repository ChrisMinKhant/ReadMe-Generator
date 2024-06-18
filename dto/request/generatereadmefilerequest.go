package dto

type GenerateReadmeFileRequest struct {
	ServiceCategory string     `json:"serviceCategory"`
	Description     string     `json:"description"`
	Endpoints       []Endpoint `json:"endpoints"`
}
