package main

import (
	"github.com/theerasakkeng/GO-CRUD/routes"
)

func main() {
	routes.InitRoute()
	routes.Router.Run()
}
