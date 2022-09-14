package scalingo

import (
	"context"
	"strings"
	"time"

	"gopkg.in/errgo.v1"
)

type InvoicesService interface {
	ListInvoices(context.Context, PaginationOpts) (Invoices, PaginationMeta, error)
	InvoiceShow(context.Context, string) (*Invoice, error)
}

var _ InvoicesService = (*Client)(nil)

const LayoutISO = "2006-01-02"

type date time.Time

type InvoiceItem struct {
	ID    string `json:"id"`
	Label string `json:"label"`
	Price int    `json:"price"`
}

type InvoiceDetailedItem struct {
	ID    string `json:"id"`
	Label string `json:"label"`
	Price int    `json:"price"`
	App   string `json:"app"`
}

type Invoice struct {
	ID                string                `json:"id"`
	TotalPrice        int                   `json:"total_price"`
	TotalPriceWithVat int                   `json:"total_price_with_vat"`
	BillingMonth      date                  `json:"billing_month"`
	PdfURL            string                `json:"pdf_url"`
	InvoiceNumber     string                `json:"invoice_number"`
	State             string                `json:"state"`
	VatRate           int                   `json:"vat_rate"`
	Items             []InvoiceItem         `json:"items"`
	DetailedItems     []InvoiceDetailedItem `json:"detailed_items"`
}

type Invoices []*Invoice

type InvoicesRes struct {
	Invoices Invoices `json:"invoices"`
	Meta     struct {
		PaginationMeta PaginationMeta `json:"pagination"`
	}
}

type InvoiceRes struct {
	Invoice *Invoice `json:"invoice"`
}

func (c *Client) ListInvoices(ctx context.Context, opts PaginationOpts) (Invoices, PaginationMeta, error) {
	var invoicesRes InvoicesRes
	err := c.ScalingoAPI().ResourceList(ctx, "account/invoices", opts.ToMap(), &invoicesRes)
	if err != nil {
		return nil, PaginationMeta{}, errgo.Mask(err)
	}
	return invoicesRes.Invoices, invoicesRes.Meta.PaginationMeta, nil
}

func (c *Client) InvoiceShow(ctx context.Context, id string) (*Invoice, error) {
	var invoiceRes InvoiceRes
	err := c.ScalingoAPI().ResourceGet(ctx, "account/invoices", id, nil, &invoiceRes)
	if err != nil {
		return nil, errgo.Mask(err)
	}
	return invoiceRes.Invoice, nil
}

func (c *date) UnmarshalJSON(b []byte) error {
	value := strings.Trim(string(b), `"`) //get rid of "
	if value == "" || value == "null" {
		return nil
	}
	if t, err := time.Parse(LayoutISO, string(value)); err != nil {
		return err
	} else {
		*c = date(t)
		return nil
	}
}

func (c date) MarshalJSON() ([]byte, error) {
	return []byte(`"` + time.Time(c).Format(LayoutISO) + `"`), nil
}
