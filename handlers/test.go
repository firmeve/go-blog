package handlers

import (
	"fmt"
	"github.com/kataras/iris"
)

func FirstTest(ctx iris.Context)  {
	fmt.Println(ctx.Application().ConfigurationReadOnly().GetOther())
	name := ctx.Params().Get("name")
	routeName := ctx.GetCurrentRoute().Name()
	ctx.Writef("Hello %s,Route name is%s", name,routeName)
	ctx.Next()
}