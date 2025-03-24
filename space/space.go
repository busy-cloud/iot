package space

import (
	"github.com/busy-cloud/boat/db"
	"github.com/busy-cloud/iot/device"
	"time"
)

func init() {
	db.Register(new(Space), new(SpaceDevice))
}

type SpaceDevice struct {
	SpaceId  string    `json:"space_id,omitempty" xorm:"pk"`
	Space    string    `json:"space,omitempty" xorm:"<-"`
	DeviceId string    `json:"device_id,omitempty" xorm:"pk"`
	Device   string    `json:"device,omitempty" xorm:"<-"`
	Created  time.Time `json:"created" xorm:"created"`
}

type Space struct {
	Id          string `json:"id" xorm:"pk"`
	Name        string `json:"name,omitempty"`        //名称
	Description string `json:"description,omitempty"` //说明

	ProjectId string `json:"project_id,omitempty" xorm:"index"`
	Project   string `json:"project,omitempty" xorm:"<-"`
	ParentId  string `json:"parent_id,omitempty" xorm:"index"`
	Parent    string `json:"parent,omitempty" xorm:"<-"`

	Disabled bool      `json:"disabled,omitempty"`
	Created  time.Time `json:"created" xorm:"created"`

	values map[string]any

	parent *Space

	devices map[string]*device.Device
}

func (s *Space) PutDevice(name string, dev *device.Device) {
	//s.devices[name] = dev
	//s.values[name] = dev.Values()
	//
	//dev.Watch(func(value map[string]any) {
	//	//此处用来触发情景模式
	//
	//})
}
