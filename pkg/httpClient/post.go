package httpClient

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"
)

func (h HttpClient) Post(ctx context.Context, path string, data interface{}, responseData interface{}) (httpStatusCode int, err error) {

	requestData, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
		return 0, err
	}

	req, err := http.NewRequest("POST", h.apiBaseURL+path, bytes.NewBuffer(requestData))
	if err != nil {
		log.Println(err)
		return 0, err
	}
	req.Header.Add("content-type", "application/json")

	resp, err := h.client.Do(req)
	if err != nil {
		log.Println(err)
		return 0, err
	}

	// if request function does not ask for response data return from here
	if responseData == nil || resp.StatusCode != http.StatusOK {
		return resp.StatusCode, nil
	}

	// if rquest function ask for response data, then parse data in requestdata interface and return
	err = h.ReadResponse(ctx, resp, responseData)
	if err != nil {
		log.Println(err)
		return resp.StatusCode, err
	}

	return resp.StatusCode, nil
}
