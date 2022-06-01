package httptools

import (
	"encoding/json"
	"errors"
	"net/http"
	"postman/internal/model"
)

func ParseRequest(httpReq *http.Request, req model.RequestDTO) error {
	jsonDecoder := json.NewDecoder(httpReq.Body)

	err := jsonDecoder.Decode(req)
	if err != nil {
		return errors.New("invalid request")
	}

	return nil
}
