package internal

import (
	"github.com/busy-cloud/boat/lib"
	"github.com/busy-cloud/iot/device"
	"time"
)

type Device struct {
	device.Device `xorm:"extends"`

	Values  map[string]any `json:"values"`
	Updated time.Time      `json:"updated"`
}

var devices lib.Map[Device]

func GetDevice(id string) *Device {
	return devices.Load(id)
}
