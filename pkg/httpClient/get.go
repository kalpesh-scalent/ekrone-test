package httpClient

import (
	"context"
	"log"
	"net/http"
)

func (h HttpClient) Get(ctx context.Context, path string, responseData interface{}) (httpStatusCode int, err error) {

	req, err := http.NewRequest("GET", h.apiBaseURL+path, nil)
	if err != nil {
		log.Println(err)
		return 0, err
	}

	resp, err := h.client.Do(req)
	if err != nil {
		log.Println(err)
		return 0, err
	}

	// if request function does not ask for response data return from here
	if responseData == nil {
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
