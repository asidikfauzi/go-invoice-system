package route

import (
	"github.com/gin-gonic/gin"
	"go-invoice-system/common/helper"
	"go-invoice-system/controller/customers"
	"go-invoice-system/controller/invoices"
	"go-invoice-system/controller/items"
	"go-invoice-system/controller/types"
)

type InitRoutes interface {
	InitRouter()
}

type RouteService struct {
	TypeService     types.TypeController         `inject:"controller_type_master"`
	CustomerService customers.CustomerController `inject:"controller_customer_master"`
	ItemService     items.ItemController         `inject:"controller_item_master"`
	InvoiceService  invoices.InvoiceController   `inject:"controller_invoice_master"`
}

func InitPackage() *RouteService {
	return &RouteService{
		TypeService:     &types.MasterTypes{},
		CustomerService: &customers.MasterCustomers{},
		ItemService:     &items.MasterItems{},
		InvoiceService:  &invoices.MasterInvoices{},
	}
}

func (r *RouteService) InitRouter() {
	router := gin.Default()

	endpoint := router.Group("/inv")
	{
		types := endpoint.Group("/type")
		{
			types.GET("", r.TypeService.GetAllTypes)
			types.GET("/:typeId", r.TypeService.FindTypeById)
			types.POST("", r.TypeService.CreateType)
			types.PATCH("/:typeId", r.TypeService.UpdateType)
			types.DELETE("/:typeId", r.TypeService.DeleteType)
		}

		customers := endpoint.Group("/customer")
		{
			customers.GET("", r.CustomerService.GetAllCustomers)
			customers.GET("/:customerId", r.CustomerService.FindCustomerById)
			customers.POST("", r.CustomerService.CreateCustomer)
			customers.PATCH("/:customerId", r.CustomerService.UpdateCustomer)
			customers.DELETE("/:customerId", r.CustomerService.DeleteCustomer)
		}

		items := endpoint.Group("/item")
		{
			items.GET("", r.ItemService.GetAllItems)
			items.GET("/:itemId", r.ItemService.FindItemById)
			items.POST("", r.ItemService.CreateItem)
			items.PATCH("/:itemId", r.ItemService.UpdateItem)
			items.DELETE("/:itemId", r.ItemService.DeleteItem)
		}

		invoices := endpoint.Group("/invoice")
		{
			invoices.GET("", r.InvoiceService.GetAllInvoices)
		}

	}

	err := router.Run(":" + helper.GetEnv("APP_PORT"))
	if err != nil {
		return
	}

}
