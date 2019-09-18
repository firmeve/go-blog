package handlers

import (
	"github.com/kataras/iris"
)

func FirstTest(ctx iris.Context)  {
	// 当遇到panic时就不会再执行了
	panic("errors")
	name := ctx.Params().Get("name")
	routeName := ctx.GetCurrentRoute().Name()
	ctx.Writef("Hello %s,Route name is%s", name,routeName)
	ctx.Next()
}