package main

import (
	"fmt"
	"github.com/blog/internal/pkg/bootstrap"
	"github.com/blog/pkg/utils"
)

func main() {
	//app := iris.Default()
	//app.OnErrorCode()
	//app.Use(recover.New())
	fmt.Println(utils.CurrentDir())
	bootstrap.RunWeb(":28184")
}