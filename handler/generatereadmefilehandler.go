package handler

import (
	"encoding/json"
	"net/http"
	"time"

	requestDto "github.com/ChrisMinKhant/megoyougo_framework/dto/request"
	responseDto "github.com/ChrisMinKhant/megoyougo_framework/dto/response"
	"github.com/ChrisMinKhant/megoyougo_framework/exception"
	"github.com/ChrisMinKhant/megoyougo_framework/service/generatereadmefileservice"
	"github.com/sirupsen/logrus"
)

type generateReadmeFileHandler struct {
	request          *requestDto.GenerateReadmeFileRequest
	exceptionHandler exception.Exception
}

func NewGenerateReadmeFileHandler() *generateReadmeFileHandler {
	return &generateReadmeFileHandler{
		request:          &requestDto.GenerateReadmeFileRequest{},
		exceptionHandler: exception.GetGeneralExceptionInstance(),
	}
}

func (generateReadmeFileHandler *generateReadmeFileHandler) Handle(w http.ResponseWriter, r *http.Request) {
	defer generateReadmeFileHandler.exceptionHandler.RecoverPanic()

	w.Header().Set("Access-Control-Allow-Origin", "*")

	err := json.NewDecoder(r.Body).Decode(generateReadmeFileHandler.request)

	if err != nil {
		logrus.Panicf("Error occurred at reading request body ::: %v\n", err.Error())
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
