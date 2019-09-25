package main

import (
	"github.com/blog/internal/pkg/bootstrap"
)

func main() {
	//app := iris.Default()
	//app.OnErrorCode()
	//app.Use(recover.New())
	bootstrap.RunWeb(":28184")
}