package middleware

import (
	"github.com/kataras/iris"
)

func Before(ctx iris.Context)  {
	ctx.Write([]byte("Global before middleware"))
	ctx.Next()
}

func After(ctx iris.Context)  {
	ctx.Write([]byte("Global after middleware"))
}

func UseGlobal(app *iris.Application)  {
	app.UseGlobal(Before)
	app.Done(After)
}