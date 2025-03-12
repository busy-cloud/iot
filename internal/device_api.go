package internal

import (
	"github.com/busy-cloud/boat/api"
	"github.com/busy-cloud/boat/curd"
	"github.com/busy-cloud/iot/types"
	"github.com/gin-gonic/gin"
)

func init() {
	api.Register("GET", "device/list", curd.ApiList[types.Device]())
	api.Register("POST", "device/create", curd.ApiCreate[types.Device]())
	api.Register("GET", "device/:id", curd.ParseParamStringId, curd.ApiGet[types.Device]())
	api.Register("POST", "device/:id", curd.ParseParamStringId, curd.ApiUpdate[types.Device]("id", "name", "description", "product_id", "disabled", "station"))
	api.Register("GET", "device/:id/delete", curd.ParseParamStringId, curd.ApiDelete[types.Device]())
	api.Register("GET", "device/:id/enable", curd.ParseParamStringId, curd.ApiDisable[types.Device](false))
	api.Register("GET", "device/:id/disable", curd.ParseParamStringId, curd.ApiDisable[types.Device](true))
	api.Register("GET", "device/:id/values", curd.ParseParamStringId, deviceValues)
}

func deviceValues(ctx *gin.Context) {
	d := devices.Load(ctx.Param("id"))
	if d == nil {
		api.Fail(ctx, "device not found")
		return
	}
	api.OK(ctx, d)
}
