package inject

import (
	"github.com/facebookgo/inject"
	"go-invoice-system/common/database"
	typeMysql "go-invoice-system/repository/mysql/types"
	"go-invoice-system/route"
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

	// SERVICE
	typeService := typeService.NewTypeService(typeMysql)

	dependencies := []*inject.Object{
		{Value: typeService, Name: "types_service"},
	}

	if liq.Routes != nil {
		dependencies = append(dependencies,
			&inject.Object{Value: liq.Routes, Name: "routes"},
			&inject.Object{Value: liq.Routes.TypeService, Name: "controller_type_master"},
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
