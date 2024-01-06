package main

import (
	"go-invoice-system/common/inject"
	"go-invoice-system/route"
)

func main() {
	routes := route.InitPackage()
	inject.DependencyInjection(inject.InjectData{
		Routes: routes,
	})

	routes.InitRouter()
}
