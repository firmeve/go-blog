package main

import (
	"github.com/blog/routes"
	"github.com/kataras/iris"
)

func main() {
	app := iris.Default()
	// Load config
	app.Configure(iris.WithConfiguration(iris.YAML("./app.yml")))
	// Load routing
	routes.Register(app)

	app.Run(iris.Addr(":28183"))
}
