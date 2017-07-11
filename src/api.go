package main

import (
	"net/http"
	"io/ioutil"
	"net/url"
	"strings"
	"strconv"
)

type Api struct {
	Name string `json:"name"`
	Url string `json:"url"`
	Methods []string `json:"methods"`
}

func NewApi(name string, url string, methods ...string) *Api {
	return &Api{name, url, methods}
}

func (api *Api) Allow(method string) bool {
	for _, m := range api.Methods {
		if m == method {
			return true
		}
	}
	return false
}

func FormatData(data map[string]interface{}) *strings.Reader {
	form := url.Values{}
	for k, v := range data {
		switch v.(type) {
		case int:
			form.Add(k, strconv.Itoa(v.(int)))
		case bool:
			form.Add(k, strconv.FormatBool(v.(bool)))
		case string:
			form.Add(k, v.(string))
		}
	}
	return strings.NewReader(form.Encode())
}

// get the api
// return status code, body, err
func (api *Api) Get() (int, string, error) {

	if !api.Allow(http.MethodGet) {
		return -1, "", NewError("not support get")
	}

	client := &http.Client{}
	req, _ := http.NewRequest(http.MethodGet, api.Url, nil)
	resp, err := client.Do(req)
	if err != nil {
		return -1, "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return -1, "", err
	}

	return resp.StatusCode, string(body), nil
}

// post the api
// return status code, body, err
func (api *Api) Post(data map[string]interface{}) (int, string, error) {

	if !api.Allow(http.MethodPost) {
		return -1, "", NewError("not support post")
	}

	client := &http.Client{}
	req, _ := http.NewRequest(http.MethodPost, api.Url, FormatData(data))
	resp, err := client.Do(req)
	if err != nil {
		return -1, "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return -1, "", err
	}

	return resp.StatusCode, string(body), nil
}

// delete the api
// return status code, body, err
func (api *Api) Delete() (int, string, error) {

	if !api.Allow(http.MethodDelete) {
		return -1, "", NewError("not support get")
	}

	client := &http.Client{}
	req, _ := http.NewRequest(http.MethodDelete, api.Url, nil)
	resp, err := client.Do(req)
	if err != nil {
		return -1, "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return -1, "", err
	}

	return resp.StatusCode, string(body), nil
}

// put the api
// return status code, body, err
func (api *Api) Put(data map[string]interface{}) (int, string, error) {

	if !api.Allow(http.MethodPut) {
		return -1, "", NewError("not support get")
	}

	client := &http.Client{}
	req, _ := http.NewRequest(http.MethodPut, api.Url, FormatData(data))
	resp, err := client.Do(req)
	if err != nil {
		return -1, "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return -1, "", err
	}

	return resp.StatusCode, string(body), nil
}



