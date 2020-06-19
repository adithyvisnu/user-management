package handler

import (
	"errors"
	"net/http"
)

// Post validates whether the request method is post or not
// return error, body and the headers
func Post(req *http.Request) (payloads map[string]interface{}, err error) {
	payloads = make(map[string]interface{}, 2)
	if req.Method != "POST" {
		return nil, errors.New("Method not allowed")
	}
	payloads["body"] = req.Body
	payloads["header"] = req.Header
	return
}

// Get validates whether the request method is get or not
// return error, params and or query strings also the header
func Get(req *http.Request) (params map[string]interface{}, err error) {
	params = make(map[string]interface{}, 2)
	if req.Method != "GET" {
		return nil, errors.New("Method not allowed")
	}
	params["query"] = req.URL.Query()
	params["header"] = req.Header
	return
}

// Put validates whether the request method is put or not
// return error, body and the headers
func Put(req *http.Request) (payloads map[string]interface{}, err error) {
	payloads = make(map[string]interface{}, 2)
	if req.Method != "POST" {
		return nil, errors.New("Method not allowed")
	}
	payloads["body"] = req.Body
	payloads["header"] = req.Header
	return
}

// Delete validates whether the request method is delete or not
// return error, params and or query strings also the header
func Delete(req *http.Request) (params map[string]interface{}, err error) {
	params = make(map[string]interface{}, 2)
	if req.Method != "DELETE" {
		return nil, errors.New("Method not allowed")
	}
	params["query"] = req.URL.Query()
	params["header"] = req.Header
	return
}
