package endpoint

import (
	"context"

	m "resuelve/invoice/invoice"

	"resuelve/api/service"

	"github.com/go-kit/kit/endpoint"
	"github.com/lnx1337/go/api"
)

/*
*@function MakeReadInvoiceEndpoint
* Middleware que recibe los parámetros GET y llama al servicio InvoiceService
* para obtener el número de facturas totales.
 */
func MakeReadInvoiceEndpoint(svc service.InvoiceService) endpoint.Endpoint {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		var (
			data interface{}
			req  *m.Invoice
			ok   bool
			err  *api.Err
		)

		// valida parámetros vacios
		if req, ok = r.(*m.Invoice); !ok {
			err = api.NewError()
			err.Push(api.Msg{Error: `MISSING_DATA`})

			return nil, err
		}

		// Se llama al servicio Invoice
		if data, err = svc.Read(*req); err.Failed() {
			return nil, err
		}

		// Respuesta para el cliente
		resp := api.Response{
			Data: data,
		}

		return resp, nil
	}
}
