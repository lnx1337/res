package request

import "fmt"

const endPoint = `http://34.209.24.195`

func Api(service string, params map[string]string) (string, error) {
	url := fmt.Sprint(endPoint, service)
	r := NewRequest(url, "GET", params)
	res, err := r.Do()
	return res, err
}
