package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	// "github.com/fir1/rest-api/pkg/erru"
)

// type ErrorResponse struct {
// 	ErrorResponse string `json:"error-message"`
// }
// type ErrArgument struct {
// 	msg string
// }

// func (e ErrArgument) Error() string {
// 	return e.msg
// }

func (s service) respond(w http.ResponseWriter, data interface{}, status int) {
	var respData interface{}
	// switch v := data.(type) {
	// case nil:
	// case erru.ErrArgument:
		 status = http.StatusBadRequest
		//  respData = ErrorResponse{ErrorMessage: v.Unwrap().Error()}
	// case error:
		 if http.StatusText(status) == "" {
			  status = http.StatusInternalServerError
		 } else {
			//   respData = ErrorResponse{ErrorMessage: v.Error()}
			status = http.StatusInternalServerError
		 }
	// default:
		 respData = data
	// }

	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	if data != nil {
		err := json.NewEncoder(w).Encode(respData)
		if err != nil {
			http.Error(w, "Could not encode in json", http.StatusBadRequest)
			return
		}
	}
}

func (s service) decode(r *http.Request, v interface{}) error {
	return json.NewDecoder(r.Body).Decode(v)
}

func (s service) readRequestBody(r *http.Request) ([]byte, error) {
	var bodyBytes []byte
	var err error
	if r.Body != nil {
		bodyBytes, err = io.ReadAll(r.Body)
		if err != nil {
			err := errors.New("could not read request body")
			return nil, err
		}
	}
	return bodyBytes, nil
}

func (s service) restoreRequestBody(r *http.Request, bodyBytes []byte) {
	r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
}