package handlers

import (
	"github.com/blog/internal/document/models"
	"github.com/blog/internal/pkg/database"
	"github.com/kataras/iris"
	"github.com/kataras/iris/versioning"
	"time"
)

func FirstTest(ctx iris.Context) {

	page := &models.Page{
		Title: "abcdefadfassd",
		Content: "abcdefafdsafdas",
	}
	page.CreatedAt = time.Now()
	database.DB.Create(page)

	//db.NewRecord(user) // => 主键为空返回`true`
	//
	//db.Create(&user)
	//
	//db.NewRecord(user) // => 创建`user`后返回`false`
	//
	//models.NewPage()
	//
	//// 当遇到panic时就不会再执行了
	//name := ctx.Params().Get("name")
	//routeName := ctx.GetCurrentRoute().Name()
	//ctx.Writef("Hello %s,Route name is%s", name, routeName)
	ctx.JSON(page)
	ctx.Next()
}

func VersionTest(ctx iris.Context) {
	ctx.WriteString(versioning.GetVersion(ctx))
	ctx.WriteString("=============version")
	ctx.Next()
}

func ResponseTest(ctx iris.Context) {
	ctx.WriteString("<br>abc<br>")
	ctx.Next()
}
