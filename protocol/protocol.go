package protocol

import (
	"github.com/busy-cloud/boat/app"
	"github.com/busy-cloud/boat/lib"
	"github.com/busy-cloud/boat/mqtt"
	"github.com/busy-cloud/boat/smart"
)

var protocols lib.Map[Protocol]

type Protocol struct {
	Name        string      `json:"name"`
	Description string      `json:"description,omitempty"`
	Station     *smart.Form `json:"station,omitempty"`
	Options     *smart.Form `json:"options,omitempty"`
}

func Register(id string, protocol *Protocol) {
	if app.Name == "" || app.Name == "boat" {
		protocols.Store(id, protocol)
	} else {
		mqtt.Publish("iot/protocol/register", protocol)
	}
}
