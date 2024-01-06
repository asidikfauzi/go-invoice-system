package main

import (
	"fmt"
	"github.com/google/uuid"
	"go-invoice-system/common/database"
	"go-invoice-system/model/domain"
	"gorm.io/gorm"
	"log"
	"time"
)

func main() {
	db, err := database.InitDatabase()
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}

	InitCustomerSeed(db)
	InitTypeSeed(db)
	InitItemSeed(db)

	fmt.Println("SUCCESSFULLY ADD SEEDER")
}

func InitCustomerSeed(db *gorm.DB) {
	customers := []domain.Customers{
		{
			IDCustomer:      uuid.New(),
			CustomerName:    "Customer Name 1",
			CustomerAddress: "Customer Address 1, City 1, Postal Code 111",
			CreatedAt:       time.Now(),
		},
		{
			IDCustomer:      uuid.New(),
			CustomerName:    "Customer Name 2",
			CustomerAddress: "Customer Address 2, City 2, Postal Code 222",
			CreatedAt:       time.Now(),
		},
		{
			IDCustomer:      uuid.New(),
			CustomerName:    "Customer Name 3",
			CustomerAddress: "Customer Address 3, City 3, Postal Code 333",
			CreatedAt:       time.Now(),
		},
	}

	for _, customer := range customers {
		var existingCustomer domain.Customers
		if err := db.Where("id_customer = ?", customer.IDCustomer).First(&existingCustomer).Error; err == nil {
			log.Printf("Customer %s already exists, skipping.", customer.IDCustomer)
			continue
		}

		if err := db.Create(&customer).Error; err != nil {
			log.Printf("Failed to create customer: %s", err)
		} else {
			log.Printf("Customer %s created successfully.", customer.IDCustomer)
		}
	}
}

func InitTypeSeed(db *gorm.DB) {
	types := []domain.Types{
		{
			IDType:    uuid.New(),
			TypeName:  "Service",
			CreatedAt: time.Now(),
		},
		{
			IDType:    uuid.New(),
			TypeName:  "Hardware",
			CreatedAt: time.Now(),
		},
	}

	for _, typ := range types {
		var existingType domain.Types
		if err := db.Where("type_name = ?", typ.TypeName).First(&existingType).Error; err == nil {
			log.Printf("Type %s already exists, skipping.", typ.TypeName)
			continue
		}

		if err := db.Create(&typ).Error; err != nil {
			log.Printf("Failed to create type: %s", err)
		} else {
			log.Printf("Type %s created successfully.", typ.TypeName)
		}
	}
}

func InitItemSeed(db *gorm.DB) {

	_, err := database.InitDatabase()
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}

	var service domain.Types
	if err = database.DB.Where("type_name = ?", "Service").First(&service).Error; err != nil {
		fmt.Println("Roles 'admin' not found:", err)
		return
	}

	var hardware domain.Types
	if err = database.DB.Where("type_name = ?", "Hardware").First(&hardware).Error; err != nil {
		fmt.Println("Roles 'admin' not found:", err)
		return
	}

	items := []domain.Items{
		{
			IDItem:       uuid.New(),
			TypeID:       service.IDType,
			ItemName:     "Design",
			ItemQuantity: 500.0,
			ItemPrice:    230.00,
			CreatedAt:    time.Now(),
		},
		{
			IDItem:       uuid.New(),
			TypeID:       service.IDType,
			ItemName:     "Development",
			ItemQuantity: 1000.0,
			ItemPrice:    100.00,
			CreatedAt:    time.Now(),
		},
		{
			IDItem:       uuid.New(),
			TypeID:       service.IDType,
			ItemName:     "Meetings",
			ItemQuantity: 1000.0,
			ItemPrice:    60.00,
			CreatedAt:    time.Now(),
		},
		{
			IDItem:       uuid.New(),
			TypeID:       hardware.IDType,
			ItemName:     "Printer",
			ItemQuantity: 30.0,
			ItemPrice:    50.00,
			CreatedAt:    time.Now(),
		},
		{
			IDItem:       uuid.New(),
			TypeID:       hardware.IDType,
			ItemName:     "Monitor",
			ItemQuantity: 20.0,
			ItemPrice:    50.00,
			CreatedAt:    time.Now(),
		},
	}

	for _, item := range items {
		var existingItem domain.Items
		if err := db.Where("item_name = ?", item.ItemName).First(&existingItem).Error; err == nil {
			log.Printf("Item %s already exists, skipping.", item.ItemName)
			continue
		}

		if err := db.Create(&item).Error; err != nil {
			log.Printf("Failed to create item: %s", err)
		} else {
			log.Printf("Item %s created successfully.", item.ItemName)
		}
	}
}
