package device

import (
	"github.com/busy-cloud/boat/api"
	"github.com/busy-cloud/boat/curd"
)

func init() {
	api.Register("GET", "device/list", curd.ApiList[Device]())
	api.Register("POST", "device/create", curd.ApiCreate[Device]())
	api.Register("GET", "device/:id", curd.ParseParamStringId, curd.ApiGet[Device]())
	api.Register("POST", "device/:id", curd.ParseParamStringId, curd.ApiUpdate[Device]("id", "name", "description", "product_id", "disabled", "station"))
	api.Register("GET", "device/:id/delete", curd.ParseParamStringId, curd.ApiDelete[Device]())
	api.Register("GET", "device/:id/enable", curd.ParseParamStringId, curd.ApiDisable[Device](false))
	api.Register("GET", "device/:id/disable", curd.ParseParamStringId, curd.ApiDisable[Device](true))
}
