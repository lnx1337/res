package request

import "testing"

func TestApiFail(t *testing.T) {
	res, err := Api("/facturas", map[string]string{})
	if err.Error() != "API_ERROR" {
		t.Error("APi error", "result: ", res, "err:", err)
	}
}

func TestApiFailInvalidParams(t *testing.T) {
	params := map[string]string{
		"id":     "717f076e-e13c-45b4-bcc4-51c229e1b326",
		"start":  "",
		"finish": "",
	}
	res, err := Api("/facturas", params)
	if err.Error() != "API_ERROR" {
		t.Error("APi error", "result: ", res, "err:", err)
	}
}

func TestApiOK(t *testing.T) {
	params := map[string]string{
		"id":     "717f076e-e13c-45b4-bcc4-51c229e1b326",
		"start":  "2017-01-01",
		"finish": "2017-01-01",
	}
	res, err := Api("/facturas", params)
	if err != nil {
		t.Error("APi error", "result: ", res, "err:", err)
	}
}
