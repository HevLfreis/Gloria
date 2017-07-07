package main

import (
	"net/http"
	"encoding/json"
	"strconv"
)

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
	json.NewEncoder(w).Encode(Response{status, msg, results})
}

func ParseForm(r *http.Request, args map[string]interface{}) (map[string]interface{}, error)  {

	data := make(map[string]interface{})

	r.ParseForm()

	for k, t := range args {

		switch t.(type) {
		case bool:
			if r.PostFormValue(k) == "" {
				data[k] = t
				break
			}

			if d, err := strconv.ParseBool(r.PostFormValue(k)); err != nil {
				return nil, err
			} else {
				data[k] = d
			}

		case int:
			if r.PostFormValue(k) == "" {
				data[k] = t
				break
			}

			if d, err := strconv.Atoi(r.PostFormValue(k)); err != nil {
				return nil, err
			} else {
				data[k] = d
			}

		case float64:
			if r.PostFormValue(k) == "" {
				data[k] = t
				break
			}

			if d, err := strconv.ParseFloat(r.PostFormValue(k), 64); err != nil {
				return nil, err
			} else {
				data[k] = d
			}

		case string:
			if r.PostFormValue(k) == "" {
				data[k] = t
				break
			}

			data[k] = r.PostFormValue(k)

		default:
			if r.PostFormValue(k) == "" {
				return nil, NewError("arg not provided: %s", k)
			}

			data[k] = r.PostFormValue(k)
		}
	}

	log.Info("parse args success", "args", data)
	return data, nil
}