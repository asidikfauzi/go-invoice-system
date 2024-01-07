package domain

import (
	"github.com/google/uuid"
	"time"
)

type Invoices struct {
	IDInvoice         uuid.UUID  `gorm:"type:char(36);unique;not null;primary_key;column:id_invoice" json:"id_invoice"`
	CustomerID        uuid.UUID  `gorm:"type:char(36);not null;" json:"customer_id"`
	InvoiceID         string     `gorm:"type:char(4);not null;" json:"invoice_id"`
	InvoiceSubject    string     `gorm:"type:varchar(255)" json:"invoice_subject"`
	InvoiceIssueDate  time.Time  `gorm:"default:null" json:"invoice_issue_date"`
	InvoiceDueDate    time.Time  `gorm:"default:null" json:"invoice_due_date"`
	InvoiceTotalItem  int        `gorm:"type:int" json:"invoice_total_item"`
	InvoiceSubTotal   float64    `gorm:"type:float" json:"invoice_sub_total"`
	InvoiceTax        float64    `gorm:"type:float" json:"invoice_tax"`
	InvoiceGrandTotal float64    `gorm:"type:float" json:"invoice_grand_total"`
	InvoiceStatus     string     `gorm:"type:varchar(10);not null;check (InvoiceStatus in ('Paid', 'Unpaid'))" json:"invoice_status"`
	CreatedAt         time.Time  `gorm:"default:null" json:"created_at"`
	CreatedByID       *uuid.UUID `gorm:"type:char(36);default:null" json:"created_by_id"`
	UpdatedAt         *time.Time `gorm:"default:null" json:"updated_at"`
	UpdatedByID       *uuid.UUID `gorm:"type:char(36);default:null" json:"updated_by_id"`
	DeletedAt         *time.Time `gorm:"default:null" json:"deleted_at"`
	DeletedByID       *uuid.UUID `gorm:"type:char(36);default:null" json:"deleted_by_id"`

	//REFERENCE
	Customers Customers `gorm:"foreignKey:CustomerID;references:id_customer" json:"customers"`
}
