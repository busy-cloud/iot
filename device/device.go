package device

import (
	"github.com/busy-cloud/boat/db"
	"time"
)

func init() {
	db.Register(&Device{})
}

type Device struct {
	Id          string         `json:"id,omitempty" xorm:"pk"`
	ProductId   string         `json:"product_id" xorm:"index"`
	Name        string         `json:"name,omitempty"`
	Description string         `json:"description,omitempty"`
	Station     map[string]any `json:"station,omitempty" xorm:"json"` //从站信息（协议定义表单）
	Disabled    bool           `json:"disabled,omitempty"`            //禁用
	Created     time.Time      `json:"created,omitempty" xorm:"created"`
}
