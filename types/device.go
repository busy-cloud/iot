package types

import "time"

type Device struct {
	Id          string    `json:"id,omitempty" xorm:"pk"`
	ProductId   string    `json:"product_id" xorm:"index"`
	Name        string    `json:"name,omitempty"`
	Description string    `json:"description,omitempty"`
	Disabled    bool      `json:"disabled,omitempty"` //禁用
	Created     time.Time `json:"created,omitempty" xorm:"created"`
}
