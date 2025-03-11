package internal

import (
	"github.com/busy-cloud/boat/api"
	"github.com/busy-cloud/boat/curd"
	"github.com/busy-cloud/boat/db"
	"github.com/busy-cloud/iot/types"
	"github.com/gin-gonic/gin"
	"io"
	"xorm.io/xorm/schemas"
)

func init() {
	api.Register("GET", "product/list", curd.ApiList[types.Product]())
	api.Register("POST", "product/create", curd.ApiCreate[types.Product]())
	api.Register("GET", "product/:id", curd.ParseParamStringId, curd.ApiGet[types.Product]())
	api.Register("POST", "product/:id", curd.ParseParamStringId, curd.ApiUpdate[types.Product]("id", "name", "description", "type", "version", "disabled"))
	api.Register("GET", "product/:id/delete", curd.ParseParamStringId, curd.ApiDelete[types.Product]())
	api.Register("GET", "product/:id/enable", curd.ParseParamStringId, curd.ApiDisable[types.Product](false))
	api.Register("GET", "product/:id/disable", curd.ParseParamStringId, curd.ApiDisable[types.Product](true))

	api.Register("GET", "product/:id/config/:config", productConfig)
	api.Register("POST", "product/:id/config/:config", productConfigUpdate)
}

func productConfig(ctx *gin.Context) {
	var config types.ProductConfig
	has, err := db.Engine().ID(schemas.PK{ctx.Param("id"), ctx.Param("config")}).Get(&config)
	if err != nil {
		api.Error(ctx, err)
		return
	}
	if !has {
		api.Fail(ctx, "找不到配置文件")
		return
	}
	api.OK(ctx, config.Content)
}

func productConfigUpdate(ctx *gin.Context) {
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		api.Error(ctx, err)
		return
	}

	config := types.ProductConfig{
		Id:      ctx.Param("id"),
		Name:    ctx.Param("name"),
		Content: string(body),
	}

	cnt, err := db.Engine().ID(schemas.PK{ctx.Param("id"), ctx.Param("config")}).Cols("content").Update(&config)
	if err != nil {
		api.Error(ctx, err)
		return
	}
	if cnt == 0 {
		api.Fail(ctx, "找不到配置文件")
		_, err = db.Engine().InsertOne(&config)
		if err != nil {
			api.Error(ctx, err)
			return
		}
	}

	api.OK(ctx, nil)
}
