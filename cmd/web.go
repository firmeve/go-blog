package main

import (
	"github.com/blog/internal/document"
	"github.com/blog/internal/pkg"
	"github.com/blog/internal/pkg/database"
	"github.com/blog/internal/pkg/http"
	"github.com/blog/internal/pkg/iris"
	iris2 "github.com/kataras/iris"
)

func main() {
	//app := iris.Default()
	//app.OnErrorCode()
	//app.Use(recover.New())
	//bootstrap.RunWeb(":28184")

	app := pkg.App()
	//
	//app.Bind(`iris`,iris2.New())
	//
	//fmt.Printf("%#v",app.Resolve(`iris`))
	//
	//os.Exit(0)
	//
	app.Register(iris.NewProvider(app))

	app.Register(
		//iris.NewProvider(app),
		http.NewProvider(app).SetBootOption(),
		database.NewProvider(app),
		document.NewProvider(app),
	)
	app.Bootstrap()

	//iris := app.Resolve(`iris`).(*iris2.Application)
	//app := http.App()
	//app.Bootstrap(http.WithProviders(
	//	database.NewServiceProvider(app),
	//	document.NewServiceProvider(app),
	//))
	app.Resolve(`iris`).(*iris2.Application).Run(iris2.Addr(":28188"))
	//app.RegisterRoutes(document.Register)
}
