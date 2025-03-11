package internal

import (
	"github.com/busy-cloud/boat/api"
	"github.com/busy-cloud/iot/protocol"
	"github.com/gin-gonic/gin"
)

func init() {

	api.Register("GET", "protocol/list", func(ctx *gin.Context) {
		var ps []*protocol.Protocol
		protocols.Range(func(name string, item *protocol.Protocol) bool {
			ps = append(ps, &protocol.Protocol{
				Name:        item.Name,
				Description: item.Description,
			})
			return true
		})
		api.OK(ctx, ps)
	})

	api.Register("GET", "protocol/:name", func(ctx *gin.Context) {
		name := protocols.Load(ctx.Param("name"))
		if name != nil {
			api.OK(ctx, name)
		} else {
			api.Fail(ctx, "协议找不到")
		}
	})

	api.Register("GET", "protocol/:name/option", func(ctx *gin.Context) {
		name := protocols.Load(ctx.Param("name"))
		if name != nil {
			api.OK(ctx, name.Options)
		} else {
			api.Fail(ctx, "协议找不到")
		}
	})

	api.Register("GET", "protocol/:name/station", func(ctx *gin.Context) {
		name := protocols.Load(ctx.Param("name"))
		if name != nil {
			api.OK(ctx, name.Station)
		} else {
			api.Fail(ctx, "协议找不到")
		}
	})
}
