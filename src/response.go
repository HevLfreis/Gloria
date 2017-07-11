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

func ParseArgs(r *http.Request, args map[string]interface{}) (map[string]interface{}, error) {

	var parse func(r *http.Request, k string) string

	switch r.Method {
	case http.MethodGet:
		parse = func(r *http.Request, k string) string {
			return r.URL.Query().Get(k)
		}
	default:
		r.ParseForm()
		parse = func(r *http.Request, k string) string {
			return r.PostFormValue(k)
		}
	}

	data := make(map[string]interface{})

	for k, t := range args {

		switch t.(type) {
		case bool:
			if parse(r, k) == "" {
				data[k] = t
				break
			}

			if d, err := strconv.ParseBool(parse(r, k)); err != nil {
				return nil, err
			} else {
				data[k] = d
			}

		case int:
			if parse(r, k) == "" {
				data[k] = t
				break
			}

			if d, err := strconv.Atoi(parse(r, k)); err != nil {
				return nil, err
			} else {
				data[k] = d
			}

		case float64:
			if parse(r, k) == "" {
				data[k] = t
				break
			}

			if d, err := strconv.ParseFloat(parse(r, k), 64); err != nil {
				return nil, err
			} else {
				data[k] = d
			}

		case string:
			if parse(r, k) == "" {
				data[k] = t
				break
			}

			data[k] = parse(r, k)

		default:
			if parse(r, k) == "" {
				return nil, NewError("arg not provided: %s", k)
			}

			data[k] = parse(r, k)
		}
	}

	log.Info("parse args success", "args", data)
	return data, nil
}