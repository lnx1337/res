package endpoint

import (
	"context"
	"net/http"
	m "resuelve/invoice"
)

/*
*@function DecodeReadInoviceRequest
* Obtiene los parámetros de la url
 */
func DecodeReadInoviceRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var inv m.Invoice

	urlParams := r.URL.Query()
	existID := len(urlParams.Get(":id")) > 0
	existStartDate := len(urlParams.Get(":startDate")) > 0
	existEndDate := len(urlParams.Get(":endDate")) > 0

	if r.ContentLength == 0 && existID && existStartDate && existEndDate {
		inv.ID = urlParams.Get(":id")
		inv.StartDate = urlParams.Get(":startDate")
		inv.EndDate = urlParams.Get(":endDate")
	}
	return &inv, nil
}
