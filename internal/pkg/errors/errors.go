package errors

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"strconv"
	_ "strconv"
)

func RegisterHandler(app *iris.Application) {
	Status404Page(app)
}

func Status404Page(app *iris.Application) {
	app.OnErrorCode(iris.StatusNotFound, func(ctx context.Context) {
		info := map[string]string{
			"status":  strconv.Itoa(iris.StatusNotFound),
			"message": "404",
		}
		ctx.View("errors/404.html", info)
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

