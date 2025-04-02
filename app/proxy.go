package app

import (
	"github.com/gin-gonic/gin"
	"strings"
)

func Proxy(ctx *gin.Context) {
	//app.zipReader = reader

	if ctx.Request.Method != "GET" {
		return
	}

	//插件 前端页面
	if str, has := strings.CutPrefix(ctx.Request.RequestURI, "/app/"); has {
		if len(str) == 0 {
			return
		}

		var app string
		var path string

		if strings.Index(str, "/") > 0 {
			app, path, _ = strings.Cut(str, "/")
		} else {
			app = str
			path = "index.html"
		}

		if p := apps.Load(app); p != nil {
			err := p.ServeFile(path, ctx)
			if err != nil {
				_ = ctx.Error(err)
			} else {
				ctx.Abort()
			}
		}
	}
}
