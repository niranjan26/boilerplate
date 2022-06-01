package responsewriter

import (
	"encoding/json"
	"net/http"
	"postman/internal/model"
)

func WriteFailResponse(resp http.ResponseWriter, code, message, version string, statusCode int) {
	response := &model.Response{
		ResponseCode:    code,
		ResponseMessage: message,
		ResponseVersion: version,
	}

	body, _ := json.Marshal(response)

	resp.WriteHeader(statusCode)
	resp.Write(body)
}

func WriteSuccessResponse(resp http.ResponseWriter, code, message, version string, data interface{}) {
	response := &model.Response{
		ResponseCode:    code,
		ResponseMessage: message,
		ResponseVersion: version,
		ResponseData:    data,
	}

	body, err := json.Marshal(response)
	if err != nil {
		WriteFailResponse(resp, "0001", "FAILED", "1", 500)
		return
	}

	resp.WriteHeader(200)
	resp.Write(body)
}
