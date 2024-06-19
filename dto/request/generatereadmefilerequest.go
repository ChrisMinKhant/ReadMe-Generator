package dto

type GenerateReadmeFileRequest struct {
	ServiceName     string     `json:"serviceName"`
	ServiceCategory string     `json:"serviceCategory"`
	Description     string     `json:"description"`
	Endpoints       []Endpoint `json:"endpoints"`
}
