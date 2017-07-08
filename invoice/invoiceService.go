package invoice

import (
	"log"
	"os"
	"resuelve/helpers"
	"resuelve/request"
	"resuelve/utils"
)

const service = `/facturas`

type InvoiceService struct {
	ID    string
	Limit int64
}

func NewInvoiceService(ID string) *InvoiceService {
	fS := &InvoiceService{}
	fS.ID = ID
	return fS
}

func (fS *InvoiceService) GetInvoice(startDate, endDate string) int64 {
	var total int64
	var numFacturas int64

	params := map[string]string{
		"id":     fS.ID,
		"start":  startDate,
		"finish": endDate,
	}

	resp, err := request.Api(service, params)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	number, isLimit := helpers.ParseResponse(resp)

	if isLimit {
		log.Print(startDate, "  a  ", endDate, "  ", resp)
		fS.Limit = number
	} else {
		log.Print(startDate, "  a  ", endDate, "  ", number)
		numFacturas = number
	}

	if numFacturas <= fS.Limit && !isLimit {
		return numFacturas
	}

	midPos := utils.GetMidDate(startDate, endDate)
	total = total + fS.GetInvoice(startDate, midPos) + fS.GetInvoice(midPos, endDate)

	return total
}
