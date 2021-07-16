package http

import (
	"github.com/asppj/droneDeploy/conf"
	"github.com/kataras/iris/v12"
)

type (
	request struct {
		Firstname string `json:"firstname"`
		Lastname  string `json:"lastname"`
	}

	response struct {
		ID       uint64 `json:"id"`       // id
		Message  string `json:"message"`  // msg
		Platform string `json:"platform"` // platform
	}
)

func Init() {
	app := iris.New()
	app.Get("/health", healthCheck)
	if err := app.Listen(":18080"); err != nil {
		panic(err)
	}
}

func healthCheck(ctx iris.Context) {
	ctx.WriteString("GET successfully\n" + conf.BuildVersion())
}
