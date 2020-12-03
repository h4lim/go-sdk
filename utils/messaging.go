package utils

import (
	"github.com/h4lim/go-sdk/app/types"
	"strconv"
	"time"
)

func GeneralMessageHandler(localize string, responseCode int) string {
	var responseMessage string
	switch localize {
	case "EN":
		responseMessage = MMEN.Entries[strconv.Itoa(responseCode)]
	case "ID":
		responseMessage = MMID.Entries[strconv.Itoa(responseCode)]
	default:
		responseMessage = MMEN.Entries[strconv.Itoa(responseCode)]
	}

	return responseMessage
}

type GeneralResponseType struct {
	ResponseStatus    bool      `json:"response_status"`
	ResponseCode      int       `json:"response_code"`
	ResponseMessage   string    `json:"response_message"`
	ResponseTimestamp time.Time `json:"response_timestamp"`
}

func GeneralResponseErrorBuild(ResponseTime time.Time, ResponseCode int, ResponseMessage string) *GeneralResponseType {
	var generalResponseType GeneralResponseType
	generalResponseType.ResponseTimestamp = ResponseTime
	generalResponseType.ResponseStatus = false
	generalResponseType.ResponseCode = ResponseCode
	generalResponseType.ResponseMessage = ResponseMessage
	return &generalResponseType
}

func GeneralResponseSuccessBuild(ResponseTime time.Time, ResponseCode int, ResponseMessage string) *GeneralResponseType {
	var generalResponseType GeneralResponseType
	generalResponseType.ResponseTimestamp = ResponseTime
	generalResponseType.ResponseStatus = true
	generalResponseType.ResponseCode = ResponseCode
	generalResponseType.ResponseMessage = ResponseMessage
	return &generalResponseType
}

func GetResponseMessage(localize string, httpCode int, responseCode int) *types.GeneralResponse {
	errorMessage := types.GeneralResponse{
		HttpCode: httpCode,
		Code:     responseCode,
		Message:  GeneralMessageHandler(localize, responseCode),
	}
	return &errorMessage
}
