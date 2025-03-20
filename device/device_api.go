package device

import (
	"github.com/busy-cloud/boat/api"
	"github.com/busy-cloud/boat/curd"
)

func init() {
	api.Register("GET", "iot/device/list", curd.ApiList[Device]())
	api.Register("POST", "iot/device/search", curd.ApiSearch[Device]())
	api.Register("POST", "iot/device/create", curd.ApiCreate[Device]())
	api.Register("GET", "iot/device/:id", curd.ApiGet[Device]())
	api.Register("POST", "iot/device/:id", curd.ApiUpdate[Device]("id", "name", "description", "product_id", "disabled", "station"))
	api.Register("GET", "iot/device/:id/delete", curd.ApiDelete[Device]())
	api.Register("GET", "iot/device/:id/enable", curd.ApiDisable[Device](false))
	api.Register("GET", "iot/device/:id/disable", curd.ApiDisable[Device](true))
}
