package main

import (
	"log"
	"resuelve/invoice"
)

func main() {

	var ID = "717f076e-e13c-45b4-bcc4-51c229e1b326"
	var startDate = "2017-01-01"
	var endDate = "2017-12-12"

	fs := invoice.NewInvoiceService(ID)
	total := fs.GetInvoice(startDate, endDate)
	log.Print("TOTAL: ", total)
}
