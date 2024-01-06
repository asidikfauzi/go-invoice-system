package inject

import (
	"fmt"
	"github.com/facebookgo/inject"
	"go-invoice-system/common/database"
	"go-invoice-system/route"
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

	fmt.Println(db)

	dependencies := []*inject.Object{}

	if liq.Routes != nil {
		dependencies = append(dependencies,
			&inject.Object{Value: liq.Routes, Name: "routes"},
		)
	}

	// dependency injection
	var g inject.Graph
	if err := g.Provide(dependencies...); err != nil {
		log.Fatal("Failed Inject Dependencies", err)
	}

	if err := g.Populate(); err != nil {
		log.Fatal("Failed Populate Inject Dependencies", err)
	}

}
