package product

import (
	"errors"
	"github.com/busy-cloud/boat/db"
	"github.com/busy-cloud/boat/lib"
	"time"
)

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
	Id         string       `json:"id,omitempty" xorm:"pk"`
	Properties []*Property  `json:"properties,omitempty" xorm:"json"`
	Events     []*Event     `json:"events,omitempty" xorm:"json"`
	Actions    []*Action    `json:"actions,omitempty" xorm:"json"`
	Validators []*Validator `json:"validators,omitempty" xorm:"json"`
	Updated    time.Time    `json:"updated,omitempty" xorm:"updated"`
	Created    time.Time    `json:"created,omitempty" xorm:"created"`
}

var modelCache = lib.Cache[ProductModel]{
	Timeout: int64(time.Minute * 10),
}

func LoadModel(id string) (*ProductModel, error) {
	m, has := modelCache.Load(id)
	if has {
		return m, nil
	}

	var pm ProductModel

	has, err := db.Engine().ID(id).Get(&pm)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New("缺少映射")
	}

	//缓存下来
	modelCache.Store(id, &pm)

	return &pm, nil
}
