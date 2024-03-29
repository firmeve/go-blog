package document

import (
	"github.com/blog/internal/document/handlers"
	"github.com/kataras/iris"
	"github.com/kataras/iris/versioning"
)

type route struct {
	iris *iris.Application
}

// Register api routes
func (r *route) registerRoutes() {
	usersAPI := r.iris.Party("/")
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

	r.iris.Get("/user/{name:string has()}", handlers.FirstTest).Name = "firstTest"
}

// Register route macros
func (r *route) registerMacros() {
	r.iris.Macros().Get("string").RegisterFunc("has", func() func(value string) bool {
		return func(value string) bool {
			return value == "abc"
		}
	})
}

func newRoute(iris *iris.Application) *route {
	return &route{
		iris: iris,
	}
}