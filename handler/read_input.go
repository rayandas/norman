package handler

import (
	"net/http"
	"io/ioutil"
	"io"
	"github.com/rancher/norman/httperror"
	"fmt"
	"encoding/json"
)

const reqMaxSize = (2 * 1<<20) + 1

var bodyMethods = map[string]bool{
	http.MethodPut: true,
	http.MethodPost: true,
}

func ReadBody(rw http.ResponseWriter, req *http.Request) (map[string]interface{}, error) {
	if !bodyMethods[req.Method] {
		return nil, nil
	}

	content, err := ioutil.ReadAll(io.LimitReader(req.Body, reqMaxSize))
	if err != nil {
		return nil, httperror.NewAPIError(httperror.INVALID_BODY_CONTENT,
			fmt.Sprintf("Body content longer than %d bytes", reqMaxSize-1))
	}

	data := map[string]interface{}{}
	if err := json.Unmarshal(content, data); err != nil {
		return nil, httperror.NewAPIError(httperror.INVALID_BODY_CONTENT,
			fmt.Sprintf("Failed to parse body: %v", err))
	}

	return data, nil
}
