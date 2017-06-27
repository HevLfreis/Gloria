package main

import "net/http"

type Response struct {
	Status  int            `json:"status"`
	Message string         `json:"message"`
	Results []interface{} `json:"results"`
}

func SetResponse(w http.ResponseWriter, resp_type string) http.ResponseWriter {
	switch resp_type {
	case "json":
		w.Header().Set("Content-Type", "application/json")
	case "html":
		w.Header().Set("Content-Type", "text/html")
	default:

	}

	w.WriteHeader(http.StatusOK)
	return w
}

func SendJsonResponse(w http.ResponseWriter, status int, msg string, results []interface{}) {
	SetResponse(w, "json")

}
