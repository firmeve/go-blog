package main

import (
	"github.com/blog/internal/document"
	"github.com/blog/internal/pkg/database"
	"github.com/blog/internal/pkg/http"
)

func main() {
	//app := iris.Default()
	//app.OnErrorCode()
	//app.Use(recover.New())
	//bootstrap.RunWeb(":28184")
	app := http.App()
	app.Bootstrap(http.WithProviders(
		database.NewServiceProvider(app),
		document.NewServiceProvider(app),
	))
	app.Run(app.ConfigFromOtherDefault(`server.addr`, `:80`).(string))
	//app.RegisterRoutes(document.Register)
}
