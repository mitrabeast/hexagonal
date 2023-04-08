package util

import (
	"encoding/json"
	"log"
	"net/http"
)

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
}

func Error(rw http.ResponseWriter, httpStatus int, err error) {
	rw.WriteHeader(httpStatus)
	resp := Response{
		Status:  "error",
		Message: err.Error(),
	}
	log.Printf("web :: error :: %s", resp.Message)
	b, err := json.Marshal(resp)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		log.Printf("web :: can't marshal response :: %v", err)
		return
	}
	if _, err = rw.Write(b); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		log.Printf("web :: can't send error to client :: %v", err)
	}
}

func Success(rw http.ResponseWriter, httpStatus int, data any) {
	rw.WriteHeader(httpStatus)
	resp := &Response{
		Status: "success",
		Data:   data,
	}
	b, err := json.Marshal(resp)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		log.Printf("web :: can't marshal response :: %v", err)
		return
	}
	if _, err = rw.Write(b); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		log.Printf("web :: can't send response to client :: %v", err)
	}
}
