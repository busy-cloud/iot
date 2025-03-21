package product

import (
	"github.com/busy-cloud/boat/db"
	"time"
)

func init() {
	db.Register(&Product{}, &ProductConfig{}, &ProductModel{})
}

type Product struct {
	Id          string    `json:"id,omitempty" xorm:"pk"`
	Name        string    `json:"name,omitempty"`
	Description string    `json:"description,omitempty"`
	Type        string    `json:"type,omitempty"` //类型
	Version     string    `json:"version,omitempty"`
	Disabled    bool      `json:"disabled,omitempty"` //禁用
	Created     time.Time `json:"created,omitempty" xorm:"created"`
}

type ProductConfig struct {
	Id      string    `json:"id,omitempty" xorm:"pk"`
	Name    string    `json:"name" xorm:"pk"` //双主键
	Content string    `json:"content,omitempty" xorm:"text"`
	Created time.Time `json:"created,omitempty" xorm:"created"`
}

// Property 属性
type Property struct {
	Name      string `json:"name,omitempty"`  //变量名称
	Label     string `json:"label,omitempty"` //显示名称
	Unit      string `json:"unit,omitempty"`  //单位
	Type      string `json:"type,omitempty"`  //bool string number array object
	Precision uint8  `json:"precision,omitempty"`
	Default   any    `json:"default,omitempty"`  //默认值
	Writable  bool   `json:"writable,omitempty"` //是否可写
	History   bool   `json:"history,omitempty"`  //是否保存历史
}

type Parameter struct {
	Name string `json:"name,omitempty"`
	Type string `json:"type,omitempty"`
}

type Event struct {
	Name        string      `json:"name,omitempty"`
	Description string      `json:"description,omitempty"`
	Parameters  []Parameter `json:"parameters,omitempty"`
}

type Action struct {
	Name        string      `json:"name,omitempty"`
	Description string      `json:"description,omitempty"`
	Parameters  []Parameter `json:"parameters,omitempty"`
	Returns     []Parameter `json:"returns,omitempty"`
}

type ProductModel struct {
	Id         string      `json:"id,omitempty" xorm:"pk"`
	Properties []*Property `json:"properties,omitempty" xorm:"json"`
	Events     []*Event    `json:"events,omitempty" xorm:"json"`
	Actions    []*Action   `json:"actions,omitempty" xorm:"json"`
	Updated    time.Time   `json:"updated,omitempty" xorm:"updated"`
	Created    time.Time   `json:"created,omitempty" xorm:"created"`
}
