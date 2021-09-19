package shared

import (
	"encoding/json"
	"net/http"
)

type WrappedResponse struct {
	Error  []string `json:"error"`
	Result []byte   `json:"result"`
}

func WriteResponse(rw http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(err.Error()))
		return
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(status)
	rw.Write(response)
}

func WriteWrappedResponse(rw http.ResponseWriter, status int, payload interface{}) {
	r, _ := json.Marshal(payload)

	wrappedResponse := WrappedResponse{
		Error:  []string{},
		Result: r,
	}
	response, err := json.Marshal(wrappedResponse)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(err.Error()))
		return
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(status)
	rw.Write(response)
}

func WriteErrorResponse(rw http.ResponseWriter, status int, errors ...string) {
	wrappedResponse := WrappedResponse{
		Error:  errors,
		Result: []byte{},
	}
	response, err := json.Marshal(wrappedResponse)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(err.Error()))
		return
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(status)
	rw.Write(response)
}
