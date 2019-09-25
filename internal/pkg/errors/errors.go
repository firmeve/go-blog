package errors

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"runtime"
	"strconv"
)

func RegisterHandler(app *iris.Application) {
	Status404Page(app)
}

func Status404Page(app *iris.Application) {
	fmt.Println(runtime.Caller(0))
	fmt.Println(runtime.Caller(1))
	fmt.Println(runtime.Caller(2))
	app.OnErrorCode(iris.StatusNotFound, func(ctx context.Context) {
		info := map[string]string{
			"status":  strconv.Itoa(iris.StatusNotFound),
			"message": "404",
		}
		ctx.View("web/errors/404.html", info)
	})
}

func Status404Json(ctx context.Context) {
	problem := iris.NewProblem().
		Status(iris.StatusNotFound).
		Key("message", "Not found")
	ctx.Problem(problem, iris.ProblemOptions{
		// Optional JSON renderer settings.
		JSON: iris.JSON{
			Indent: "  ",
		},
	})
}
