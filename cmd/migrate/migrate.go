package main

import (
	"fmt"
	"go-invoice-system/common/database"
	"go-invoice-system/model/domain"
)

func main() {
	_, err := database.InitDatabase()
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}

	err = database.DB.AutoMigrate(&domain.Customers{})
	if err != nil {
		panic("Error Create Database Customers")
	}

	err = database.DB.AutoMigrate(&domain.Types{})
	if err != nil {
		panic("Error Create Database Types")
	}

	err = database.DB.AutoMigrate(&domain.Items{})
	if err != nil {
		panic("Error Create Database Items")
	}

	err = database.DB.AutoMigrate(&domain.Invoices{})
	if err != nil {
		panic("Error Create Database Invoice")
	}

	err = database.DB.AutoMigrate(&domain.InvoiceHasItems{})
	if err != nil {
		panic("Error Create Database InvoiceHasItems")
	}

	fmt.Println("SUCCESSFULLY ADD MIGRATION")
}
