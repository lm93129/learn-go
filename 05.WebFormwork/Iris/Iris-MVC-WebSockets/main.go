package main

import (
	"sync/atomic"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/websocket"
)

func main() {
	app := iris.New()
	// 加载模板
	app.RegisterView(iris.HTML("./views", ".html"))

	// 渲染 ./views/index.html.
	app.Get("/", func(ctx iris.Context) {
		ctx.View("index.html")
	})

	mvc.Configure(app.Party("/websocket"), configureMVC)
	// Or mvc.New(app.Party(...)).Configure(configureMVC)

	// http://localhost:8080
	app.Run(iris.Addr(":8080"))
}

func configureMVC(m *mvc.Application) {
	ws := websocket.New(websocket.Config{})
	// http://localhost:8080/websocket/iris-ws.js
	m.Router.Any("/iris-ws.js", websocket.ClientHandler())

	// 将会绑定 ws.Upgrade 的结果 ，它是一个  websocket.Connection
	// 由 `m.Handle` 服务的控制器.
	m.Register(ws.Upgrade)
	m.Handle(new(websocketController))
}

var visits uint64

func increment() uint64 {
	return atomic.AddUint64(&visits, 1)
}

func decrement() uint64 {
	return atomic.AddUint64(&visits, ^uint64(0))
}

type websocketController struct {
	// 注意：你也可以使用匿名字段，因为 binder 将会找到它
	//
	// 这是当前 websocket 的连接,每个客户端都有自己的  *websocketController 实例对象。
	Conn websocket.Connection
}

func (c *websocketController) onLeave(roomName string) {
	// visits--
	newCount := decrement()
	// 这将在所有客户端上调用 「visit」 事件，除了当前的客户端，除了当前的这个，
	// （当前的不能访问时因为，它已经断开了，但是对于任何情况都是使用这种类型的设计）
	_ = c.Conn.To(websocket.Broadcast).Emit("visit", newCount)
}

func (c *websocketController) update() {
	// visits++
	newCount := increment()

	// 这将在所有客户端上调用 「visit」 事件，包括当前的 websocket
	// 使用 'newCount' variable 变量。
	//
	// 你有很多方法可以做得更快，例如你可以发送一个新 vistor
	// 且客户端可以自行增加，但在这里我们只是「展示」websocket 控制器
	c.Conn.To(websocket.All).Emit("visit", newCount)
}

func (c *websocketController) Get( /* websocket.Connection could be lived here as well, it doesn't matter */ ) {
	c.Conn.OnLeave(c.onLeave)
	c.Conn.On("visit", c.update)

	// call it after all event callbacks registration.
	c.Conn.Wait()
}
