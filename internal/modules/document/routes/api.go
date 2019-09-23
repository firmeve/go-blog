package routes

import (
	"github.com/blog/internal/handlers"
	"github.com/blog/internal/middleware"
	"github.com/kataras/iris"
	"github.com/kataras/iris/versioning"
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

	usersAPI := app.Party("/")
	// version 1.
	usersAPIV1 := versioning.NewGroup(">= 1, < 2")
	usersAPIV1.Get("/api/users", func(ctx iris.Context) {
		ctx.Writef("v1 resource: /api/users handler")
	})
	usersAPIV1.Post("/new", func(ctx iris.Context) {
		ctx.Writef("v1 resource: /api/users/new post handler")
	})

	usersAPI.Get("api/response",handlers.ResponseTest)

	// version 2.
	usersAPIV2 := versioning.NewGroup(">= 2, < 3")
	usersAPIV2.Get("/api/users", handlers.VersionTest)
	usersAPIV2.Post("/", handlers.VersionTest)

	versioning.RegisterGroups(usersAPI, versioning.NotFoundHandler, usersAPIV1,usersAPIV2)

	app.Get("/user/{name:string has()}", handlers.FirstTest).Name = "firstTest"
}
