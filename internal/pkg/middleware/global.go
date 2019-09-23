package middleware

import (
	"fmt"
	"github.com/kataras/iris"
	//"strings"
)

func Before(ctx iris.Context)  {
	//ctx.Values().Set(versioning.Key, ctx.URLParamDefault("version", "2.0"))
	//ctx.Write([]byte("Global before middleware"))

	//record response
	ctx.Record()
	ctx.Next()
}

func After(ctx iris.Context)  {
	// Record Body
	fmt.Println(string(ctx.Recorder().Body()))
	//fmt.Println("After")
	//fmt.Println("zzzzz")
	ctx.Write([]byte("Global after middleware"))
}

func UseGlobal(app *iris.Application)  {
	app.UseGlobal(Error,Before)
	app.DoneGlobal(After)
	//app.Done(After)
}

func Error(ctx iris.Context)  {
	defer func() {
		if err := recover(); err != nil {
			ctx.StatusCode(500)
			ctx.Text( err.(string))
			ctx.StopExecution()
			//fmt.Println(err)
			//if httpError,ok := err.(error); ok {
			//	fmt.Println("GGGGGGGGGGGGG")
			//	//if strings.ToLower(ctx.GetHeader("Accept")) == "application/json" {
			//		ctx.Text("error" + httpError.Error())
			//	//}
			//	//ctx.StopExecution()
			//}
		}


	}()
	//
	ctx.Next()
}