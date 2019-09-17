package routes

import (
	"github.com/blog/handlers"
	"github.com/blog/middleware"
	"github.com/kataras/iris"
)

func Register(app *iris.Application) {
	registerMacros(app)
	registerRoutes(app)
}

// Register route macros
func registerMacros(app *iris.Application) {
	app.Macros().Get("string").RegisterFunc("has", func() func(value string) bool {
		return func(value string) bool {
			return value == "abc"
		}
	})
}

// Register api routes
func registerRoutes(app *iris.Application) {
	// Register global middleware
	middleware.UseGlobal(app)

	app.Get("/user/{name:string has()}", handlers.FirstTest).Name = "firstTest"
}
