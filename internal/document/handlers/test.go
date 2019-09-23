package handlers

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/versioning"
)

func FirstTest(ctx iris.Context)  {
	// 当遇到panic时就不会再执行了
	name := ctx.Params().Get("name")
	routeName := ctx.GetCurrentRoute().Name()
	ctx.Writef("Hello %s,Route name is%s", name,routeName)
	ctx.Next()
}

func VersionTest(ctx iris.Context)  {
	ctx.WriteString(versioning.GetVersion(ctx))
	ctx.WriteString("=============version")
	ctx.Next()
}

func ResponseTest(ctx iris.Context)  {
	ctx.WriteString("<br>abc<br>")
	ctx.Next()
}