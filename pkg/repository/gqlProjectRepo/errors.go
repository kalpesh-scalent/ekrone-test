package gqlProjectRepo

import "errors"

var (
	errQueryNotFound     = errors.New("queryNotFound")
	errWrongResponseCode = errors.New("wrong query response status code")
	errResponse          = errors.New("error response from graphQL request")
)
