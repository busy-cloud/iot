package protocol

import (
	"github.com/busy-cloud/boat/app"
	"github.com/busy-cloud/boat/lib"
	"github.com/busy-cloud/boat/mqtt"
	"github.com/busy-cloud/boat/smart"
)

var protocols lib.Map[Protocol]

type Base struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	Version     string `json:"version,omitempty"`
	Author      string `json:"author,omitempty"`
	Copyright   string `json:"copyright,omitempty"`
}

type Protocol struct {
	Base

	Station *smart.Form `json:"station,omitempty"`
	Options *smart.Form `json:"options,omitempty"`
}

func Register(protocol *Protocol) {
	if app.Name == "" || app.Name == "boat" {
		protocols.Store(protocol.Name, protocol)
	} else {
		mqtt.Publish("iot/register/protocol", protocol)
	}
}
