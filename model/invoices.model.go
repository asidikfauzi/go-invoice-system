package model

import (
	"github.com/google/uuid"
	"time"
)

type (
	Invoices struct {
		IDInvoice         uuid.UUID `json:"id_invoice"`
		CustomerID        uuid.UUID `json:"customer_id"`
		InvoiceID         string    `json:"invoice_id"`
		InvoiceSubject    string    `json:"invoice_subject"`
		InvoiceIssueDate  time.Time `json:"invoice_issue_date"`
		InvoiceDueDate    time.Time `json:"invoice_due_date"`
		InvoiceTotalItem  int       `json:"invoice_total_item"`
		InvoiceSubTotal   float64   `json:"invoice_sub_total"`
		InvoiceTax        float64   `json:"invoice_tax"`
		InvoiceGrandTotal float64   `json:"invoice_grand_total"`
		InvoiceStatus     string    `json:"invoice_status"`
		Customer          Customers `gorm:"foreignKey:CustomerID;" json:"customers"`
	}

	GetInvoices struct {
		IDInvoice        uuid.UUID `json:"id_invoice"`
		InvoiceID        string    `json:"invoice_id"`
		InvoiceSubject   string    `json:"invoice_subject"`
		InvoiceIssueDate time.Time `json:"invoice_issue_date"`
		InvoiceDueDate   time.Time `json:"invoice_due_date"`
		InvoiceTotalItem int       `json:"invoice_total_item"`
		InvoiceStatus    string    `json:"invoice_status"`
		CustomerName     string    `json:"customer_name"`
	}

	GetInvoice struct {
		IDInvoice         uuid.UUID   `json:"id_invoice"`
		InvoiceID         string      `json:"invoice_id"`
		InvoiceSubject    string      `json:"invoice_subject"`
		InvoiceIssueDate  time.Time   `json:"invoice_issue_date"`
		InvoiceDueDate    time.Time   `json:"invoice_due_date"`
		InvoiceTotalItem  int         `json:"invoice_total_item"`
		InvoiceSubTotal   float64     `json:"invoice_sub_total"`
		InvoiceTax        float64     `json:"invoice_tax"`
		InvoiceGrandTotal float64     `json:"invoice_grand_total"`
		InvoiceStatus     string      `json:"invoice_status"`
		Customer          GetCustomer `json:"customer"`
		Items             []GetItem   `json:"items"`
	}

	RequestInvoices struct {
		InvoiceID        string `json:"invoice_id"`
		InvoiceSubject   string `json:"invoice_subject"`
		InvoiceIssueDate string `json:"invoice_issue_date"`
		InvoiceDueDate   string `json:"invoice_due_date"`
		InvoiceTotalItem int    `json:"invoice_total_item"`
		InvoiceStatus    string `json:"invoice_status"`
		CustomerName     string `json:"customer_name"`
	}

	RequestInvoice struct {
		InvoiceID         string          `json:"invoice_id" validate:"required,max:4,min:4"`
		InvoiceSubject    string          `json:"invoice_subject" validate:"required"`
		InvoiceIssueDate  string          `json:"invoice_issue_date" validate:"required,date"`
		InvoiceDueDate    string          `json:"invoice_due_date" validate:"required,date"`
		InvoiceTotalItem  int             `json:"invoice_total_item" validate:"required,number"`
		InvoiceSubTotal   float64         `json:"invoice_sub_total" validate:"required,number"`
		InvoiceTax        float64         `json:"invoice_tax" validate:"required,number"`
		InvoiceGrandTotal float64         `json:"invoice_grand_total" validate:"required,number"`
		InvoiceStatus     string          `json:"invoice_status" validate:"required"`
		Customer          RequestCustomer `json:"customer" validate:"required"`
		Items             []RequestItem   `json:"items" validate:"required"`
	}
)
