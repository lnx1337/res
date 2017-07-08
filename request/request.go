package request

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// API ERRORS
const apiError = `API_ERROR`

// Request tiene las propiedades publicas
type Request struct {
	Url    string
	Params map[string]string
	Method string
}

// NewRequest permite crear un objeto tipo request para hacer peticiones GET
func NewRequest(url string, method string, params map[string]string) *Request {
	r := &Request{}
	r.Url = url
	r.Method = method
	r.Params = params
	return r
}

// Do Ejecuta la petición y parsea el response
func (r *Request) Do() (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest(r.Method, r.Url, nil)

	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	q := req.URL.Query()

	for k, v := range r.Params {
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
