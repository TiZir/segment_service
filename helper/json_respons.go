package helper

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func MakeRespons(w http.ResponseWriter, message string, code int, err error) {

	if err != nil {
		message = fmt.Sprintf("%s: %s", message, err.Error())
	}

	Response := Response{
		Message: message,
		Code:    code,
	}

	response, err := json.Marshal(Response)
	if err != nil {
		log.Println("Failed to marshal error response:", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
