package bootstrap

import (
	"github.com/blog/internal/document"
	"github.com/blog/internal/pkg/errors"
	"github.com/blog/internal/pkg/routes"
	"github.com/kataras/iris"
	"sync"
)

var (
	webApp *iris.Application
	webOnce sync.Once
)

func Web() *iris.Application  {
	if webApp != nil{
		return webApp
	}

	webOnce.Do(func() {
		webApp = iris.New()
		// Load config
		webApp.Configure(iris.WithConfiguration(iris.YAML("config/app.yml")))
		errors.RegisterHandler(webApp)
		// Load routing
		routes.RegisterRoutes(webApp,document.Register)

		webApp.Get("/product-problem", newProductProblemRender)
	})

	return webApp
}

func RunWeb(addr string)  {
	Web().Run(iris.Addr(addr))
}

func newProductProblemRender(ctx iris.Context)  {
	ctx.Problem(newProductProblem("abc","ef"),iris.ProblemOptions{
		//JSON: iris.JSON{
		//},
	})
}

func newProductProblem(productName, detail string) iris.Problem {
	return iris.NewProblem().
		// The type URI, if relative it automatically convert to absolute.
		Type("/product-error2").
		// The title, if empty then it gets it from the status code.
		//Title("Product validation problem").
		// Any optional details.
		//Detail(detail).
		// The status error code, required.
		Status(iris.StatusBadRequest).
		// Any custom key-value pair.
		Key("message", detail)
	// Optional cause of the problem, chain of Problems.
	// .Cause(other iris.Problem)
}

func problemExample(ctx iris.Context) {
	/*
		p := iris.NewProblem().
			Type("/validation-error").
			Title("Your request parameters didn't validate").
			Detail("Optional details about the error.").
			Status(iris.StatusBadRequest).
		 	Key("customField1", customValue1)
		 	Key("customField2", customValue2)
		ctx.Problem(p)
		// OR
		ctx.Problem(iris.Problem{
			"type":   "/validation-error",
			"title":  "Your request parameters didn't validate",
			"detail": "Optional details about the error.",
			"status": iris.StatusBadRequest,
		 	"customField1": customValue1,
		 	"customField2": customValue2,
		})
		// OR
	*/

	// Response like JSON but with indent of "  " and
	// content type of "application/problem+json"
	ctx.Problem(newProductProblem("product name", "problem error details"), iris.ProblemOptions{
		// Optional JSON renderer settings.
		//JSON: iris.JSON{
		//	Indent: "  ",
		//},
		// OR
		// Render as XML:
		// RenderXML: true,
		// XML:       iris.XML{Indent: "  "},
		//
		// Sets the "Retry-After" response header.
		//
		// Can accept:
		// time.Time for HTTP-Date,
		// time.Duration, int64, float64, int for seconds
		// or string for date or duration.
		// Examples:
		// time.Now().Add(5 * time.Minute),
		// 300 * time.Second,
		// "5m",
		//
		RetryAfter: 300,
		// A function that, if specified, can dynamically set
		// retry-after based on the request. Useful for ProblemOptions reusability.
		// Overrides the RetryAfter field.
		//
		// RetryAfterFunc: func(iris.Context) interface{} { [...] }
	})
}
