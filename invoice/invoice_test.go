package invoice

import "testing"

func TestObjInovice(t *testing.T) {
	id := "717f076e-e13c-45b4-bcc4-51c229e1b326"
	inv := NewInvoice(id)

	if inv.ID != id {
		t.Error("Expected", id)
	}
}

func TestGetInvoiceFail(t *testing.T) {
	id := "717f076e-e13c-45b4-bcc4-51c229e1b326"
	inv := NewInvoice(id)
	total, _ := inv.GetInvoice("2017-01-02", "2017-02-10")
	if total == 0 {
		t.Error("Expected:", 0, "total:", total)
	}
}

func TestGetInvoiceOK(t *testing.T) {
	id := "717f076e-e13c-45b4-bcc4-51c229e1b326"
	inv := NewInvoice(id)
	total, _ := inv.GetInvoice("2017-01-01", "2017-12-12")
	if total > 152 {
		t.Error("Expected:", 0, "total:", total)
	}
}
