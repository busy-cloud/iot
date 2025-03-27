package internal

import (
	"github.com/busy-cloud/boat/db"
	"github.com/busy-cloud/boat/lib"
	"github.com/busy-cloud/boat/mqtt"
	"github.com/busy-cloud/iot/device"
	"github.com/busy-cloud/iot/product"
	"github.com/busy-cloud/iot/project"
	"github.com/busy-cloud/iot/space"
	"time"
)

var devices lib.Map[Device]

func GetDevice(id string) *Device {
	return devices.Load(id)
}

type Device struct {
	device.Device `xorm:"extends"`

	Values  map[string]any `json:"values"`
	Updated time.Time      `json:"updated"`

	projects []string
	spaces   []string

	model *product.ProductModel
}

func (d *Device) Open() error {

	//查询绑定的项目
	var ps []*project.ProjectDevice
	err := db.Engine().Where("device_id=?", d.Id).Find(&ps) //.Distinct("project_id")
	if err != nil {
		return err
	}
	for _, p := range ps {
		d.projects = append(d.projects, p.ProjectId)
	}

	//查询绑定的设备
	var ss []*space.SpaceDevice
	err = db.Engine().Where("device_id=?", d.Id).Find(&ss) //.Distinct("space_id")
	if err != nil {
		return err
	}
	for _, s := range ss {
		d.spaces = append(d.spaces, s.SpaceId)
	}

	//加载物模型
	d.model, err = product.LoadModel(d.ProductId)
	if err != nil {
		return err
	}

	return nil
}

func (d *Device) PutValues(values map[string]any) {
	d.Values = values
	d.Updated = time.Now()

	//广播消息
	var topics []string
	for _, p := range d.projects {
		topics = append(topics, "project/"+p+"/device/"+d.Id+"/property")
	}
	for _, s := range d.spaces {
		topics = append(topics, "space/"+s+"/device/"+d.Id+"/property")
	}
	if len(topics) > 0 {
		mqtt.PublishEx(topics, values)
	}
}
