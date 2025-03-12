package product

import (
	"github.com/busy-cloud/boat/api"
	"github.com/busy-cloud/boat/curd"
	"github.com/busy-cloud/boat/db"
	"github.com/gin-gonic/gin"
	"io"
	"xorm.io/xorm/schemas"
)

func init() {
	api.Register("GET", "product/list", curd.ApiList[Product]())
	api.Register("POST", "product/create", curd.ApiCreate[Product]())
	api.Register("GET", "product/:id", curd.ParseParamStringId, curd.ApiGet[Product]())
	api.Register("POST", "product/:id", curd.ParseParamStringId, curd.ApiUpdate[Product]("id", "name", "description", "type", "version", "protocol", "disabled"))
	api.Register("GET", "product/:id/delete", curd.ParseParamStringId, curd.ApiDelete[Product]())
	api.Register("GET", "product/:id/enable", curd.ParseParamStringId, curd.ApiDisable[Product](false))
	api.Register("GET", "product/:id/disable", curd.ParseParamStringId, curd.ApiDisable[Product](true))

	//物模型
	api.Register("GET", "product/:id/model", curd.ApiGet[ProductModel]())
	api.Register("POST", "product/:id/model", curd.ApiUpdate[ProductModel]("properties", "events", "actions"))

	//配置接口，一般用于协议点表等
	api.Register("GET", "product/:id/config/:config", productConfig)
	api.Register("POST", "product/:id/config/:config", productConfigUpdate)
}

func productConfig(ctx *gin.Context) {
	var config ProductConfig
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

	config := ProductConfig{
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
