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
	bodyByte, _ := json.Marshal(body)

	req, err := http.NewRequest(method, url, bytes.NewBuffer(bodyByte))
	if err != nil {
		return 0, err
	}

	resp, err := http.DefaultClient.Do(req)
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
