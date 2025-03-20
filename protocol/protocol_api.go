package protocol

import (
	"github.com/busy-cloud/boat/api"
	"github.com/gin-gonic/gin"
)

func init() {

	api.Register("GET", "iot/protocol/list", func(ctx *gin.Context) {
		var ps []*Base
		protocols.Range(func(name string, item *Protocol) bool {
			ps = append(ps, &item.Base)
			return true
		})
		api.OK(ctx, ps)
	})

	api.Register("GET", "iot/protocol/:name", func(ctx *gin.Context) {
		p := protocols.Load(ctx.Param("name"))
		if p != nil {
			api.OK(ctx, p)
		} else {
			api.Fail(ctx, "协议找不到")
		}
	})

	api.Register("GET", "iot/protocol/:name/option", func(ctx *gin.Context) {
		p := protocols.Load(ctx.Param("name"))
		if p != nil {
			api.OK(ctx, p.Options)
		} else {
			api.Fail(ctx, "协议找不到")
		}
	})

	api.Register("GET", "iot/protocol/:name/station", func(ctx *gin.Context) {
		p := protocols.Load(ctx.Param("name"))
		if p != nil {
			api.OK(ctx, p.Station)
		} else {
			api.Fail(ctx, "协议找不到")
		}
	})
}
