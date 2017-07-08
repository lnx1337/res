package request

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const apiError = `API_ERROR`

type Request struct {
	Url    string
	Params map[string]string
	Method string
}

func NewRequest(url string, method string, params map[string]string) *Request {
	r := &Request{}
	r.Url = url
	r.Method = method
	r.Params = params
	return r
}

func (self *Request) Do() (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest(self.Method, self.Url, nil)

	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	q := req.URL.Query()

	for k, v := range self.Params {
		q.Add(k, v)
	}

	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)

	if err != nil {
		return "", errors.New(apiError)
	}

	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		return string(respBody), errors.New(apiError)
	}

	return string(respBody), err
}
