package service

import (
	m "resuelve/invoice/invoice"

	"github.com/lnx1337/go/api"
)

type (
	InvoiceSvc     struct{}
	InvoiceService interface {
		Read(m.Invoice) (interface{}, *api.Err)
	}
	Response struct {
		Total            int64 `json:"total"`
		NumberOfRequests int64 `json:"numberOfRequests"`
	}
)

/*
*@function Read
* Se encarga de llamar al modulo para realizar la búsqueda de facturas.
* @var i m.Invoice son los parámetros de la URL ID, StartDate, EndDate.
 */
func (InvoiceSvc) Read(i m.Invoice) (interface{}, *api.Err) {
	// Se crea una instancia al modulo de búsqueda.
	inv := m.NewInvoice(i.ID)
	total, err := inv.GetInvoice(i.StartDate, i.EndDate)

	// Manejo de errores
	if err != nil {
		jserr := api.NewError()
		jserr.Push(api.Msg{Error: err.Error()})
		return nil, jserr
	}

	// En caso de tener éxito
	resp := Response{
		Total:            total,
		NumberOfRequests: inv.TotalRequests,
	}

	return resp, &api.Err{}
}
