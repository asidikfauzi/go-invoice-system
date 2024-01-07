package inject

import (
	"github.com/facebookgo/inject"
	"go-invoice-system/common/database"
	customerMysql "go-invoice-system/repository/mysql/customers"
	typeMysql "go-invoice-system/repository/mysql/types"
	"go-invoice-system/route"
	customerService "go-invoice-system/service/customers"
	typeService "go-invoice-system/service/types"
	"log"
)

type InjectData struct {
	Routes *route.RouteService
}

func DependencyInjection(liq InjectData) {
	db, err := database.InitDatabase()
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}

	// MYSQL
	typeMysql := typeMysql.NewTypeMysql(db)
	customerMysql := customerMysql.NewCustomerMysql(db)

	// SERVICE
	typeService := typeService.NewTypeService(typeMysql)
	customerService := customerService.NewCustomerService(customerMysql)

	dependencies := []*inject.Object{
		{Value: typeService, Name: "types_service"},
		{Value: customerService, Name: "customers_service"},
	}

	if liq.Routes != nil {
		dependencies = append(dependencies,
			&inject.Object{Value: liq.Routes, Name: "routes"},
			&inject.Object{Value: liq.Routes.TypeService, Name: "controller_type_master"},
			&inject.Object{Value: liq.Routes.CustomerService, Name: "controller_customer_master"},
		)
	}

	// dependency injection
	var g inject.Graph
	if err = g.Provide(dependencies...); err != nil {
		log.Fatal("Failed Inject Dependencies", err)
	}

	if err = g.Populate(); err != nil {
		log.Fatal("Failed Populate Inject Dependencies", err)
	}

}
