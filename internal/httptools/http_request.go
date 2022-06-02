package httptools

import (
	"context"
	"errors"
	"log"
	"net/http"
	"strings"

	"bytes"
	"encoding/json"
	"fmt"
	"io"

	"github.com/corsc/go-commons/iocloser"
)

func MakeHTTPRequest(ctx context.Context, client http.Client, url, method string, requestBody interface{}, headers, queryParams map[string]string) (int, []byte, error) {
	body, _ := json.Marshal(requestBody)

	url = addQueryParams(url, queryParams)

	httpReq, err := http.NewRequestWithContext(context.Background(), method, url, bytes.NewBuffer(body))
	if err != nil {
		log.Print(fmt.Sprintf("unable to create request %s", err))
		return 0, nil, errors.New("bad request")
	}

	for key, value := range headers {
		httpReq.Header.Set(key, value)
	}

	resp, err := client.Do(httpReq)

	defer RespBodyClose(resp)

	if err != nil {
		return 0, nil, errors.New(fmt.Sprint("failed to call api with err: %s", err))
	}

	if resp == nil {
		return 0, nil, errors.New("no response from http call")
	}

	payload, err := io.ReadAll(resp.Body)
	if err != nil {
		return resp.StatusCode, nil, errors.New(fmt.Sprint("failed to read response err: %s", err))
	}

	return resp.StatusCode, payload, nil
}

func addQueryParams(url string, queryParams map[string]string) (resp string) {
	resp = url

	if queryParams != nil {
		var pairs []string
		for key, value := range queryParams {
			pairs = append(pairs, fmt.Sprint(key, "=", value))
		}

		resp = url + "?" + strings.Join(pairs, "&")
	}

	return resp
}

func RespBodyClose(resp *http.Response) {
	if resp != nil {
		iocloser.Close(resp.Body)
	}
}
