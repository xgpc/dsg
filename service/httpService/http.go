package httpService

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
)

const ReqGet = "GET"
const ReqPost = "POST"

func DoRequest(url string, method string, bodyContent []byte) ([]byte, error) {
	client := &http.Client{}
	request := &http.Request{}
	var err error
	if method == ReqGet {
		request, err = http.NewRequest(ReqGet, url, nil)
		if err != nil {
			return nil, err
		}
	} else {
		request, err = http.NewRequest(ReqPost, url, bytes.NewReader(bodyContent))
		if err != nil {
			return nil, err
		}
	}

	request.Header.Set("Content-Type", "application/json")

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, nil
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(response.Body)

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}
	return body, err
}
