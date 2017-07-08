package invoice

import (
	"log"
	"os"
	"resuelve/helpers"
	"resuelve/request"
	"resuelve/utils"
)

const service = `/facturas`

// @struct invoice.InvoiceService.
// @field ID de factura.
// @field limit usado para obtener el limite de resultados
// que devuelve el servidor.
type InvoiceService struct {
	ID    string
	Limit int64
}

// @function NewInvoiceService.
// @returns retorna un objeto tipo InvoiceService.
// @var ID obtiene el id de la factura.
func NewInvoiceService(ID string) *InvoiceService {
	fS := &InvoiceService{}
	fS.ID = ID
	return fS
}

/*
* @function GetInvoice permite hacer peticiones en el servidor utilizando
* un algoritmo de busqueda recursivo
* @var startDate, @var endDate rango de fechas para obtener el numero
* de facturas en ese rango
 */
func (fS *InvoiceService) GetInvoice(startDate, endDate string) int64 {
	var total int64
	var numFacturas int64

	// Parametros para enviar al servidor.
	params := map[string]string{
		"id":     fS.ID,
		"start":  startDate,
		"finish": endDate,
	}
	// Se envia la petición al servidor.
	resp, err := request.Api(service, params)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
	// Se interpreta la respuesta.
	number, isLimit := helpers.ParseResponse(resp)

	if isLimit {
		log.Print(startDate, "  a  ", endDate, "  ", resp)
		fS.Limit = number
	} else {
		log.Print(startDate, "  a  ", endDate, "  ", number)
		numFacturas = number
	}

	// Algoritmo de busqueda

	// Se verifica si el numero de facturas recibido por el servidor
	// es diferente de el limite, si es menor igual al limite entonces
	// retorna el valor para sumarlo al total de facturas.
	if numFacturas <= fS.Limit && !isLimit {
		return numFacturas
	}

	// Se calcula la fecha intermedia entre dos fechas.
	midPos := utils.GetMidDate(startDate, endDate)
	// Inicia recursividad.
	total = total + fS.GetInvoice(startDate, midPos) + fS.GetInvoice(midPos, endDate)

	return total
}
