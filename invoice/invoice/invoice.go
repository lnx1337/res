package invoice

import (
	"resuelve/invoice/helpers"
	"resuelve/invoice/request"
	"resuelve/invoice/utils"
)

const service = `/facturas`

/*
* @struct invoice.Invoice.
* @field ID de factura.
* @field limit usado para obtener el límite de resultados
* que devuelve el servidor.
 */
type Invoice struct {
	ID            string `json:"id"`
	StartDate     string `json:"startDate"`
	EndDate       string `json:"endDate"`
	Limit         int64  `json:"-"`
	TotalRequests int64  `json:"-"`
}

/*
* @function NewInvoice.
* @returns retorna un objeto tipo Invoice.
* @var ID obtiene el id de la factura.
 */
func NewInvoice(ID string) *Invoice {
	fS := &Invoice{}
	fS.ID = ID
	return fS
}

/*
* @function GetInvoice permite hacer peticiones en el servidor utilizando
* un algoritmo de búsqueda recursivo
* @var startDate, @var endDate rango de fechas para obtener el número
* de facturas en ese rango.
 */
func (fS *Invoice) GetInvoice(startDate, endDate string) (int64, error) {
	var total int64
	var numFacturas int64

	// parámetros para enviar al servidor.
	params := map[string]string{
		"id":     fS.ID,
		"start":  startDate,
		"finish": endDate,
	}
	// Se envia la petición al servidor.
	resp, err := request.Api(service, params)
	if err != nil {
		return 0, err
	}

	fS.TotalRequests++

	// Se interpreta la respuesta.
	number, isLimit := helpers.ParseResponse(resp)

	if isLimit {
		//	log.Print("Limite:  ", startDate, "  a  ", endDate, "  ", resp)
		fS.Limit = number
	} else {
		//	log.Print(startDate, "  a  ", endDate, "   Encontradas: ", number)
		numFacturas = number
	}

	// Algoritmo de busqueda

	// Se verifica si el número de facturas recibido por el servidor
	// es diferente de el limite, si es menor igual al limite entonces
	// retorna el valor para sumarlo al total de facturas.
	if numFacturas <= fS.Limit && !isLimit {
		return numFacturas, nil
	} else if !isLimit {
		return numFacturas, nil
	}

	// Se calcula la fecha intermedia entre dos fechas.
	midPos := utils.GetMidDate(startDate, endDate)

	// Se suma un día a la fecha del segundo calculo para evitar
	// sumar facturas del día ya contado
	datePlus := utils.Parse(midPos).AddDate(0, 0, 1)
	dateMidPlus := utils.GetDate(datePlus)

	// Inicia recursividad.
	resp1, _ := fS.GetInvoice(startDate, midPos)
	resp2, _ := fS.GetInvoice(dateMidPlus, endDate)

	total = total + resp1 + resp2

	return total, nil
}
