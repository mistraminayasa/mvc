package main

import (
	"mvc/routers"

	"github.com/kataras/iris"
)

func main() {

	r := routers.GetEngine()
	r.Run(iris.Addr(":8080"))

}
