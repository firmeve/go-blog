package main

import (
	"github.com/blog/internal/document"
	"github.com/blog/internal/pkg/http"
)

func main() {
	//app := iris.Default()
	//app.OnErrorCode()
	//app.Use(recover.New())
	//bootstrap.RunWeb(":28184")
	app := http.App()
	app.Default()
	app.RegisterRoutes(document.Register)
	app.Run(":28183")
}