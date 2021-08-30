package http

import (
	"encoding/json"
	"net/http"
)

func Response(w http.ResponseWriter, r *http.Request, v interface{}, statusCode int) {
	_, ok := headers["Content-Type"]

	if !ok {
		SetHeader("Content-Type", "application/json")
	}

	for key, value := range headers {
		w.Header().Add(key, value)
	}

	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(v)

	ResetHeaders()
}
