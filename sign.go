package perseus

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type perseusResponse struct {
	Timestamp string `json:"x-khronos"`
	Signature string `json:"x-gorgon"`
}

var PerseusURL = "http://localhost:4123/"

func Sign(url string, query string) (string, string, error) {
	payload := []byte(
		fmt.Sprintf(
			`{"url":"%s","query":"%s"}`,
			url,
			query,
		),
	)

	request, err := http.NewRequest(
		"POST",
		PerseusURL,
		bytes.NewBuffer(payload),
	)
	if err != nil {
		return "", "", err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Content-Length", strconv.Itoa(len(payload)))

	client := http.Client{
		Timeout: time.Second * 5,
	}

	response, err := client.Do(request)
	if err != nil {
		return "", "", err
	}

	responseRaw, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", "", err
	}

	perseusResponse := perseusResponse{}

	err = json.Unmarshal(responseRaw, &perseusResponse)
	if err != nil {
		return "", "", err
	}

	return perseusResponse.Timestamp, perseusResponse.Signature, nil
}
