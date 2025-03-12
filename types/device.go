package types

import (
	"github.com/busy-cloud/boat/db"
	"time"
)

func init() {
	db.Register(&Device{})
}

type Device struct {
	Id          string    `json:"id,omitempty" xorm:"pk"`
	ProductId   string    `json:"product_id" xorm:"index"`
	Name        string    `json:"name,omitempty"`
	Description string    `json:"description,omitempty"`
	Station     string    `json:"station,omitempty"`  //从站信息（协议定义表单）
	Disabled    bool      `json:"disabled,omitempty"` //禁用
	Created     time.Time `json:"created,omitempty" xorm:"created"`
}
