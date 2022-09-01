package driver

import (
	"bytes"
	"echo-get-started/types"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

func doPost(uri string, reqBody types.JSON) (types.JSON, error) {
	jsonStr, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	res, err := http.Post(uri, "application/json", bytes.NewBuffer(jsonStr))
	if err != nil {
		return nil, err
	}

	if res.StatusCode >= 400 {
		return nil, &types.HttpError{StatusCode: res.StatusCode}
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	s := string(body)
	if s == "" {
		return types.JSON{}, nil
	}

	var value types.JSON

	return value, nil
}

func doGet(uri string) ([]byte, error) {
	request, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}

	if res.StatusCode >= 400 {
		return nil, &types.HttpError{StatusCode: res.StatusCode}
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	return body, nil
}

func doGetWithQuery(uri string, query_params url.Values) ([]byte, error) {
	request, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return nil, err
	}

	request.URL.RawQuery = query_params.Encode()

	res, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}

	if res.StatusCode >= 400 {
		return nil, &types.HttpError{StatusCode: res.StatusCode}
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	return body, nil
}

func doDelete(uri string) error {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodDelete, uri, nil)
	if err != nil {
		return err
	}

	res, err := client.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode >= 400 {
		return &types.HttpError{StatusCode: res.StatusCode}
	}

	defer res.Body.Close()

	return nil
}

func doPut(uri string, body types.JSON) (types.JSON, error) {
	jsonString, err := json.Marshal(body)

	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPut, uri, bytes.NewBuffer(jsonString))
	if err != nil {
		return nil, err
	}

	req.Header.Set("content-type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode >= 400 {
		return nil, &types.HttpError{StatusCode: res.StatusCode}
	}

	defer res.Body.Close()
	responseBody, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	if string(responseBody) == "" {
		return types.JSON{}, nil
	}

	var value types.JSON
	err = json.Unmarshal([]byte(responseBody), &value)
	if err != nil {
		return nil, err
	}

	return value, nil
}

func doPatch(uri string, body types.JSON) (types.JSON, error) {
	jsonString, err := json.Marshal(body)

	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPatch, uri, bytes.NewBuffer(jsonString))

	if err != nil {
		return nil, err
	}

	req.Header.Set("content-type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode >= 400 {
		return nil, &types.HttpError{StatusCode: res.StatusCode}
	}

	defer res.Body.Close()
	responseBody, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	if string(responseBody) == "" {
		return types.JSON{}, nil
	}

	var value types.JSON
	err = json.Unmarshal([]byte(responseBody), &value)
	if err != nil {
		return nil, err
	}

	return value, nil
}
