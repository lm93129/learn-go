package main

import (
	"github.com/kataras/iris"
)

func main() {
	app := iris.New()
	app.OnErrorCode(iris.StatusNotFound, notFound)
	app.OnErrorCode(iris.StatusInternalServerError, internalServerError)
	// 为所有的大于等于400的状态码注册一个处理器：
	// app.OnAnyErrorCode(handler)
	app.Get("/", func(ctx iris.Context) {
		_ = ctx.View("index.html")
	})
	_ = app.Run(iris.Addr(":8080"))
}

func notFound(ctx iris.Context) {
	// 出现 404 的时候，就跳转到 $views_dir/errors/404.html 模板
	_ = ctx.View("errors/404.html")
}

func internalServerError(ctx iris.Context) {
	_, _ = ctx.WriteString("Oups something went wrong, try again")
}
