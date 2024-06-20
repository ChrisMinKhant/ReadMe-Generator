package handler

import (
	"encoding/json"
	"net/http"
	"time"

	requestDto "github.com/ChrisMinKhant/megoyougo_framework/dto/request"
	responseDto "github.com/ChrisMinKhant/megoyougo_framework/dto/response"
	"github.com/ChrisMinKhant/megoyougo_framework/service/generatereadmefileservice"
)

type generateReadmeFileHandler struct {
	request *requestDto.GenerateReadmeFileRequest
}

func NewGenerateReadmeFileHandler() *generateReadmeFileHandler {
	return &generateReadmeFileHandler{
		request: &requestDto.GenerateReadmeFileRequest{},
	}
}

func (generateReadmeFileHandler *generateReadmeFileHandler) Handle(w http.ResponseWriter, r *http.Request) {
	err := json.NewDecoder(r.Body).Decode(generateReadmeFileHandler.request)

	if err != nil {
		panic("Error occurred at reading request body ::: " + err.Error())
	}

	service := generatereadmefileservice.New()

	w.Header().Set("Content-Type", "application/json")

	if service.GenerateReadmeFile(generateReadmeFileHandler.request) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&responseDto.GenereteReadmeFileResponse{
			Status:    "200",
			Message:   "Successful",
			Timestamp: time.Now(),
		})

		return
	}

	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(&responseDto.GenereteReadmeFileResponse{
		Status:    "500",
		Message:   "Fail",
		Timestamp: time.Now(),
	})
}
