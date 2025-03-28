package device

import (
	"github.com/busy-cloud/boat/db"
	"github.com/busy-cloud/iot/product"
	"time"
)

func init() {
	db.Register(&Device{}, &DeviceModel{})
}

type Device struct {
	Id          string         `json:"id,omitempty" xorm:"pk"`
	ProductId   string         `json:"product_id,omitempty" xorm:"index"`
	LinkerId    string         `json:"linker_id,omitempty" xorm:"index"`
	IncomingId  string         `json:"incoming_id,omitempty" xorm:"index"`
	Name        string         `json:"name,omitempty"`
	Description string         `json:"description,omitempty"`
	Station     map[string]any `json:"station,omitempty" xorm:"json"` //从站信息（协议定义表单）
	Disabled    bool           `json:"disabled,omitempty"`            //禁用
	Created     time.Time      `json:"created,omitempty" xorm:"created"`
}

type DeviceModel struct {
	Id         string               `json:"id,omitempty" xorm:"pk"`
	Validators []*product.Validator `json:"validators,omitempty" xorm:"json"`
	Created    time.Time            `json:"created,omitempty" xorm:"created"`
}
