package routes

import (
	"github.com/kataras/iris"
)

type Routes func(app *iris.Application)

//func Register(app *iris.Application) {
//	registerMacros(app)
//	registerRoutes(app)
//}
//
//// Register route macros
//func registerMacros(app *iris.Application) {
//	app.Macros().Get("string").RegisterFunc("has", func() func(value string) bool {
//		return func(value string) bool {
//			return value == "abc"
//		}
//	})
//}

// Register api routes
func RegisterRoutes(app *iris.Application, routes ...Routes) {
	//Register global middleware
	//middleware.UseGlobal(app)

	for _, route := range routes {
		route(app)
	}
}
