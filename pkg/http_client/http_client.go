package httpclient

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func HitAPI(
	method string,
	url string,
	header map[string]string,
	body interface{},
	result interface{},
) (
	int,
	error,
) {

	//TODO ADD METRIX
	payloadBuf := new(bytes.Buffer)
	err := json.NewEncoder(payloadBuf).Encode(body)
	if err != nil {
		return 0, err
	}

	req, err := http.NewRequest(method, url, payloadBuf)
	if err != nil {
		return 0, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}

	defer resp.Body.Close()

	bodyResponse, _ := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(bodyResponse, &result)
	if err != nil {
		return 0, err
	}

	return resp.StatusCode, nil
}
