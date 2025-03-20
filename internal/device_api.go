package internal

import (
	"github.com/busy-cloud/boat/api"
	"github.com/busy-cloud/boat/curd"
	"github.com/gin-gonic/gin"
)

func init() {
	api.Register("GET", "iot/device/:id/values", curd.ParseParamStringId, deviceValues)
}

func deviceValues(ctx *gin.Context) {
	d := devices.Load(ctx.Param("id"))
	if d == nil {
		api.Fail(ctx, "device not found")
		return
	}
	api.OK(ctx, d)
}
