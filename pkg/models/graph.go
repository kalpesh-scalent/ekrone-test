package models

import "encoding/json"

type GraphRequest struct {
	Query string `json:"query"`
}

type Error map[string]interface{}

type GraphResponse struct {
	Data   json.RawMessage `json:"data"`
	Errors []Error         `json:"errors"`
}
