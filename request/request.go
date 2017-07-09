package request

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// @var apiError constante de errores
const apiError = `API_ERROR`

/*
* @struct Request
* @field URL endpoint
* @Params parámetros para enviar en URL
* @Method tipo de petición GET, POST
 */
type Request struct {
	URL    string
	Params map[string]string
	Method string
}

/*
* @fuction NewRequest Crea un objeto tipo request para hacer peticiones
* @var url endpoint
* @var method GET, POST
* @var params parámetros para enviar en la URL
 */
func NewRequest(url string, method string, params map[string]string) *Request {
	r := &Request{}
	r.URL = url
	r.Method = method
	r.Params = params
	return r
}

// @function Do Ejecuta la petición y parsea la respuesta.
func (r *Request) Do() (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest(r.Method, r.URL, nil)

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
	// se destruye el objeto
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		return string(respBody), errors.New(apiError)
	}

	return string(respBody), err
}
